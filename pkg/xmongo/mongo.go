package xmongo

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func New(dbName string, opts ...Option) (*Service, error) {
	conf, err := getConfig(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to load config for mongodb: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	clientOpts := options.Client()
	clientOpts.ApplyURI(conf.MongoDBURI)
	clientOpts.SetMonitor(otelmongo.NewMonitor())

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %w", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping mongodb: %w", err)
	}

	return &Service{
		Database: client.Database(dbName),
		Client:   client,
	}, nil
}
