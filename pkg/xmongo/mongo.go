package xmongo

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"

	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	lgr      logr.Logger
	Client   *mongo.Client
	Database *mongo.Database
}

func New(ctx context.Context, lgr logr.Logger, dbName string, opts ...Option) (*Service, error) {
	conf, err := getConfig(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to load config for mongodb: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	clientOpts := options.Client()
	clientOpts.ApplyURI(conf.MongoDBURI)
	clientOpts.SetMonitor(otelmongo.NewMonitor())

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %w", err)
	}

	ctx, cancel = context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping mongodb: %w", err)
	}

	return &Service{
		lgr:      lgr.WithName("xmongo"),
		Client:   client,
		Database: client.Database(dbName),
	}, nil
}

func (svc *Service) Stop(ctx context.Context) error {
	svc.lgr.Info("disconnecting from the database")
	return svc.Client.Disconnect(ctx)
}
