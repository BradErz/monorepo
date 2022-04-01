package xgrpc

import (
	"net"
	"time"

	"google.golang.org/grpc"
)

type serverConfig struct {
	gracePeriod   time.Duration
	grpcOpts      []grpc.ServerOption
	listener      net.Listener
	registerFuncs []RegisterServerFunc
}

type RegisterServerFunc func(s *grpc.Server)

type ServerOption func(conf *serverConfig)

func WithGracePeriod(dur time.Duration) ServerOption {
	return func(conf *serverConfig) {
		conf.gracePeriod = dur
	}
}

func WithGRPCServerOption(opt grpc.ServerOption) ServerOption {
	return func(conf *serverConfig) {
		conf.grpcOpts = append(conf.grpcOpts, opt)
	}
}

func WithListener(lis net.Listener) ServerOption {
	return func(conf *serverConfig) {
		conf.listener = lis
	}
}

func WithRegisterFunc(fn RegisterServerFunc) ServerOption {
	return func(conf *serverConfig) {
		conf.registerFuncs = append(conf.registerFuncs, fn)
	}
}

func getServerConfig(opts ...ServerOption) *serverConfig {
	conf := &serverConfig{}
	for _, opt := range opts {
		opt(conf)
	}

	return conf
}
