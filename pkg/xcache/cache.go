package xcache

import (
	"fmt"

	"github.com/go-redis/redis/extra/redisotel/v9"
	"github.com/go-redis/redis/v9"
)

type Client struct {
	Client redis.UniversalClient
	config *Config
}

func New(opts ...Option) (*Client, error) {
	conf, err := getConfig(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to load config for mongodb: %w", err)
	}

	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: conf.RedisAddrs,
	})

	if err := redisotel.InstrumentTracing(rdb); err != nil {
		return nil, fmt.Errorf("failed to add tracing to redis client: %w", err)
	}
	if err := redisotel.InstrumentMetrics(rdb); err != nil {
		return nil, fmt.Errorf("failed to add metrics to redis client: %w", err)
	}

	return &Client{
		Client: rdb,
		config: conf,
	}, nil
}

func (client *Client) Prefix(key string) string {
	return client.config.Namespace + ":" + key
}
