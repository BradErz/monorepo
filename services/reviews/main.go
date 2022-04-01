package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/BradErz/monorepo/pkg/xlogger"

	"github.com/BradErz/monorepo/pkg/telemetry"

	"github.com/BradErz/monorepo/pkg/xgrpc"

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
	lgr, err := xlogger.New()
	if err != nil {
		return fmt.Errorf("failed to create xlogger: %w", err)
	}

	if _, e := telemetry.Init(lgr, telemetry.WithServiceName("reviews")); e != nil {
		return fmt.Errorf("failed to setup telemetry: %w", e)
	}

	mon, err := xmongo.New(lgr, "reviews-service")
	if err != nil {
		return fmt.Errorf("failed to create mongoclient: %w", err)
	}
	defer func() {
		_ = mon.Stop(context.Background())
	}()
	store := storage.NewReviews(mon.Database)

	svc := service.NewReviews(store)
	reviewsSrv, err := web.New(lgr, svc)
	if err != nil {
		return fmt.Errorf("failed to listen on port: %w", err)
	}

	grpcSrv, err := xgrpc.NewServer(lgr,
		xgrpc.WithGracePeriod(time.Second*2),
		xgrpc.WithRegisterFunc(web.Register(reviewsSrv)),
	)
	if err != nil {
		return fmt.Errorf("failed to create grpc.Server: %w", err)
	}

	if err := grpcSrv.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}

	return nil
}
