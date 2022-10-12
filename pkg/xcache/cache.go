package xcache

import (
	"fmt"

	"github.com/go-redis/redis/v9"
)

type Cache struct {
	Client redis.UniversalClient
}

func New(opts ...Option) (*Cache, error) {
	conf, err := getConfig(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to load config for mongodb: %w", err)
	}

	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: conf.RedisAddrs,
	})
	return &Cache{
		Client: rdb,
	}, nil
}
