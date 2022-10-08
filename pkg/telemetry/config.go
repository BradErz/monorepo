package telemetry

import "github.com/kelseyhightower/envconfig"

type Config struct {
	ServiceName             string `envconfig:"SERVICE_NAME"`
	Environment             string `envconfig:"ENV"`
	Enabled                 bool   `envconfig:"ENABLED"`
	MetricsEnabled          bool
	JaegerCollectorEndpoint string `envconfig:"JAEGER_COLLECTOR_ENDPOINT"`
}

type Option func(conf *Config)

func WithEnabled() Option {
	return func(conf *Config) {
		conf.Enabled = true
	}
}

func WithMetricsEnabled() Option {
	return func(conf *Config) {
		conf.MetricsEnabled = true
	}
}

func WithServiceName(s string) Option {
	return func(conf *Config) {
		conf.ServiceName = s
	}
}

func WithJaegerCollectorEndpoint(u string) Option {
	return func(conf *Config) {
		conf.JaegerCollectorEndpoint = u
	}
}

func defaultConfig() (*Config, error) {
	conf := &Config{
		Enabled:                 false,
		Environment:             "dev",
		JaegerCollectorEndpoint: "http://localhost:14268/api/traces",
	}
	return conf, envconfig.Process("TELEMETRY", conf)
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
