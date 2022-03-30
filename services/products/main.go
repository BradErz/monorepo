package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BradErz/monorepo/pkg/telemetry"
	"github.com/oklog/run"

	"github.com/BradErz/monorepo/pkg/xgrpc"

	"github.com/BradErz/monorepo/pkg/xmongo"

	"github.com/BradErz/monorepo/services/products/storage"

	"github.com/BradErz/monorepo/services/products/service"
	"github.com/BradErz/monorepo/services/products/web"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := app(); err != nil {
		logrus.WithError(err).Fatal("failed to start application")
	}
}

func app() error {
	lgr := logrus.NewEntry(logrus.New())

	if err := telemetry.Init(lgr, telemetry.WithServiceName("products")); err != nil {
		return fmt.Errorf("failed to setup telemetry: %w", err)
	}

	mon, err := xmongo.New("products-service")
	if err != nil {
		return fmt.Errorf("failed to create mongoclient: %w", err)
	}
	defer mon.Stop(context.Background())

	store, err := storage.NewProducts(mon.Database)
	if err != nil {
		return fmt.Errorf("failed to connect to mongodb: %w", err)
	}

	svc := service.NewProducts(store)
	productsSrv, err := web.New(lgr, svc)
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
				lgr.WithError(err).Error("failed to shutdown")
			}
		})
	}

	err = g.Run()
	if err == context.Canceled {
		return nil
	}

	return err
}
