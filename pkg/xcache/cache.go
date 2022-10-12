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

	rdb.AddHook(redisotel.NewTracingHook())

	return &Client{
		Client: rdb,
		config: conf,
	}, nil
}

func (client *Client) Prefix(key string) string {
	return client.config.Namespace + ":" + key
}
