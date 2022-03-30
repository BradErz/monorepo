package main

import (
	"context"
	"fmt"
	"time"

	"github.com/BradErz/monorepo/pkg/telemetry"

	"github.com/BradErz/monorepo/pkg/xgrpc"

	"github.com/BradErz/monorepo/pkg/xmongo"

	"github.com/BradErz/monorepo/services/reviews/storage"

	"github.com/BradErz/monorepo/services/reviews/service"
	"github.com/BradErz/monorepo/services/reviews/web"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := app(); err != nil {
		logrus.WithError(err).Fatal("failed to start application")
	}
}

func app() error {
	lgr := logrus.NewEntry(logrus.New())

	if err := telemetry.Init(lgr, telemetry.WithServiceName("reviews")); err != nil {
		return fmt.Errorf("failed to setup telemetry: %w", err)
	}
	mon, err := xmongo.New("reviews-service")
	if err != nil {
		return fmt.Errorf("failed to create mongoclient: %w", err)
	}
	defer mon.Stop(context.Background())

	store, err := storage.NewReviews(mon.Database)
	if err != nil {
		return fmt.Errorf("failed to connect to mongodb: %w", err)
	}

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
