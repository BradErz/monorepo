package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BradErz/monorepo/gen/go/products/v1/productsv1connect"
	"github.com/BradErz/monorepo/gen/go/reviews/v1/reviewsv1connect"
	"github.com/BradErz/monorepo/pkg/xconnect"
	"github.com/BradErz/monorepo/pkg/xlogger"
	"github.com/bufbuild/connect-go"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/BradErz/monorepo/pkg/telemetry"

	"github.com/BradErz/monorepo/pkg/xmongo"

	"github.com/BradErz/monorepo/services/reviews/storage"

	"github.com/BradErz/monorepo/services/reviews/service"
	"github.com/BradErz/monorepo/services/reviews/web"
)

func main() {
	if err := app(); err != nil {
		log.Fatalf("failed to start application: %s", err)
	}
}

func app() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	lgr, err := xlogger.New()
	if err != nil {
		return fmt.Errorf("failed to create xlogger: %w", err)
	}

	_, tpShutdown, err := telemetry.Init(lgr, telemetry.WithServiceName("reviews"), telemetry.WithMetricsEnabled())
	if err != nil {
		return fmt.Errorf("failed to setup telemetry: %w", err)
	}

	db, err := xmongo.New(ctx, lgr, "reviews-service")
	if err != nil {
		return fmt.Errorf("failed to create mongoclient: %w", err)
	}

	productsV1Client := productsv1connect.NewProductsServiceClient(
		&http.Client{
			Timeout: time.Second,
			Transport: otelhttp.NewTransport(http.DefaultTransport,
				// operation is always sent as an empty sting...
				otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
					return "http.client " + r.URL.Path
				}),
			),
		},
		os.Getenv("API_PRODUCTS_V1_URL"),
	)

	store := storage.NewReviews(db.Database)

	svc := service.NewReviews(store, productsV1Client)
	reviewsSrv := web.New(lgr, svc)

	mux := http.NewServeMux()
	interceptors := connect.WithInterceptors(
		xconnect.ErrorsInterceptor(),
		xconnect.LogrInterceptor(lgr),
	)
	mux.Handle(reviewsv1connect.NewReviewsServiceHandler(reviewsSrv, interceptors))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to create listener: %w", err)
	}
	srv := xconnect.NewServer(mux, lis)

	go func() {
		if err := srv.Start(); err != nil {
			lgr.Error(err, "failed to start server")
		}
	}()
	lgr.Info("application started", "addr", srv.Addr())

	// wait for shutdown
	<-ctx.Done()
	lgr.Info("application shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		lgr.Error(err, "failed to shutdown server")
	}

	if err := db.Stop(ctx); err != nil {
		lgr.Error(err, "failed to stop database connections")
	}

	tpShutdown(ctx)

	return nil
}
