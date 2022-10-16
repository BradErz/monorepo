package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/BradErz/monorepo/pkg/xcache"
	"github.com/BradErz/monorepo/pkg/xconnect"
	"github.com/BradErz/monorepo/pkg/xlogger"
	"github.com/bufbuild/connect-go"

	"github.com/BradErz/monorepo/pkg/telemetry"
	productsv1connect "go.buf.build/bufbuild/connect-go/braderz/monorepo/products/v1/productsv1connect"

	"github.com/BradErz/monorepo/pkg/xmongo"

	"github.com/BradErz/monorepo/services/cache"
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
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	lgr, err := xlogger.New()
	if err != nil {
		return fmt.Errorf("failed to create xlogger: %w", err)
	}

	_, tpShutdown, err := telemetry.Init(lgr, telemetry.WithServiceName("products"), telemetry.WithMetricsEnabled())
	if err != nil {
		return fmt.Errorf("failed to setup telemetry: %w", err)
	}

	db, err := xmongo.New(ctx, lgr, "products-service")
	if err != nil {
		return fmt.Errorf("failed to create mongo client: %w", err)
	}
	cacheClient, err := xcache.New(xcache.WithNamespace("products-service"))
	if err != nil {
		return fmt.Errorf("failed to create cache client: %w", err)
	}

	svc := service.NewProducts(lgr,
		storage.NewProducts(db.Database),
		cache.NewProducts(cacheClient),
	)

	productsSrv := web.New(lgr, svc)

	mux := http.NewServeMux()
	interceptors := connect.WithInterceptors(
		xconnect.ErrorsInterceptor(),
		xconnect.LogrInterceptor(lgr),
	)
	mux.Handle(productsv1connect.NewProductsServiceHandler(productsSrv, interceptors))

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
