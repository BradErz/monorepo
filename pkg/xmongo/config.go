package xmongo

import "github.com/kelseyhightower/envconfig"

type Config struct {
	MongoDBURI string `envconfig:"MONGODB_URI"`
}

type Option func(config *Config)

func WithURI(uri string) Option {
	return func(config *Config) {
		config.MongoDBURI = uri
	}
}

func defaultConfig() (*Config, error) {
	conf := &Config{
		MongoDBURI: "mongodb://localhost:27017",
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
