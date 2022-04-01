package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/BradErz/monorepo/pkg/xlogger"

	reviewsv1 "github.com/BradErz/monorepo/gen/go/reviews/v1"
	"github.com/BradErz/monorepo/pkg/telemetry"
	"github.com/oklog/run"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"

	"github.com/BradErz/monorepo/pkg/xgrpc"

	"github.com/BradErz/monorepo/pkg/xmongo"

	"github.com/BradErz/monorepo/services/products/storage"

	"github.com/BradErz/monorepo/services/products/service"
	"github.com/BradErz/monorepo/services/products/web"
)

func main() {
	if err := app(); err != nil {
		log.Fatalf("failed to start application: %s", err)
	}
}

func app() error {
	lgr, err := xlogger.New()
	if err != nil {
		return fmt.Errorf("failed to create xlogger: %w", err)
	}

	if _, err := telemetry.Init(lgr, telemetry.WithServiceName("products")); err != nil {
		return fmt.Errorf("failed to setup telemetry: %w", err)
	}

	mon, err := xmongo.New(lgr, "products-service")
	if err != nil {
		return fmt.Errorf("failed to create mongoclient: %w", err)
	}
	defer mon.Stop(context.Background())

	store := storage.NewProducts(mon.Database)

	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	}
	reviewsConn, err := grpc.Dial("reviews:50051", grpcOpts...)
	if err != nil {
		return fmt.Errorf("failed to dial to reviews service: %w", err)
	}
	reviewsClient := reviewsv1.NewReviewsServiceClient(reviewsConn)

	svc := service.NewProducts(store)
	productsSrv, err := web.New(lgr, svc, reviewsClient)
	if err != nil {
		return fmt.Errorf("failed to listen on port: %w", err)
	}

	grpcSrv, err := xgrpc.NewServer(lgr,
		xgrpc.WithGracePeriod(time.Second*2),
		xgrpc.WithRegisterFunc(web.Register(productsSrv)),
	)
	if err != nil {
		return fmt.Errorf("failed to create grpc.Server: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	var g run.Group
	{
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case <-c:
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		}, func(e error) {
			cancel()
		})
	}

	{
		g.Add(func() error {
			return grpcSrv.ListenAndServe()
		}, func(err error) {
			if err := grpcSrv.Shutdown(err); err != nil {
				lgr.Error(err, "failed to shutdown")
			}
		})
	}

	err = g.Run()
	if err == context.Canceled {
		return nil
	}

	return err
}
