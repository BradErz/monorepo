package xcache

import "github.com/kelseyhightower/envconfig"

type Config struct {
	RedisAddrs []string `envconfig:"REDIS_ADDRS"`
	Namespace  string   `envconfig:"NAMESPACE"`
}

type Option func(config *Config)

func WithAddr(addrs ...string) Option {
	return func(config *Config) {
		config.RedisAddrs = addrs
	}
}

func WithNamespace(namespace string) Option {
	return func(config *Config) {
		config.Namespace = namespace
	}
}

func defaultConfig() (*Config, error) {
	conf := &Config{
		RedisAddrs: []string{"localhost:6379"},
		Namespace:  "",
	}
	return conf, envconfig.Process("", conf)
}

func getConfig(opts ...Option) (*Config, error) {
	conf, err := defaultConfig()
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(conf)
	}

	return conf, nil
}
