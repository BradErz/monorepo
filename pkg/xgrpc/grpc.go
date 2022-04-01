package xgrpc

import (
	"context"
	"fmt"
	"net"

	"github.com/go-logr/logr"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"github.com/BradErz/monorepo/pkg/xerrors"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	lgr        logr.Logger
	conf       *serverConfig
	grpcServer *grpc.Server
}

func NewServer(lgr logr.Logger, opts ...ServerOption) (*Server, error) {
	conf, err := getServerConfig(opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to get serverConfig: %w", err)
	}

	// define some standard options we want to use
	grpcOpts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			otelgrpc.StreamServerInterceptor(),
			grpc_recovery.StreamServerInterceptor(),
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			// grpc_logrus.StreamServerInterceptor(lgr),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			otelgrpc.UnaryServerInterceptor(),
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			// grpc_logrus.UnaryServerInterceptor(lgr),
			grpc_recovery.UnaryServerInterceptor(),
			ErrorMapping,
		)),
	}

	// merge the options
	grpcOpts = append(grpcOpts, conf.grpcOpts...)
	grpcSrv := grpc.NewServer(grpcOpts...)

	for _, registerFunc := range conf.registerFuncs {
		registerFunc(grpcSrv)
	}

	reflection.Register(grpcSrv)

	return &Server{
		lgr:        lgr.WithName("xgrpcserver"),
		grpcServer: grpcSrv,
		conf:       conf,
	}, nil
}

func (srv *Server) ListenAndServe() error {
	if srv.conf.listener == nil {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			return fmt.Errorf("failed to create listner")
		}
		srv.conf.listener = lis
	}
	srv.lgr.Info("starting grpc server", "addr", srv.conf.listener.Addr())
	return srv.grpcServer.Serve(srv.conf.listener)
}

func (srv *Server) Shutdown(err error) error {
	srv.lgr.Info("grpc server: shutting down")
	if srv.conf.gracePeriod == 0 {
		srv.grpcServer.Stop()
		srv.lgr.Info("grpc server: shutdown  without grace period")
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), srv.conf.gracePeriod)
	defer cancel()

	stopped := make(chan struct{})
	go func() {
		srv.lgr.Info("waiting before stopping", "period", srv.conf.gracePeriod)
		srv.grpcServer.GracefulStop()
		close(stopped)
	}()

	select {
	case <-ctx.Done():
		srv.grpcServer.Stop()
		return nil
	case <-stopped:
		cancel()
	}

	srv.lgr.Info("successfully shutdown")
	return nil
}

// exampleAuthFunc is used by a middleware to authenticate requests
func exampleAuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	if token != "ciccio" {
		return nil, xerrors.Newf(xerrors.CodeUnauthenticated, "invalid auth token: %v", err)
	}

	return context.WithValue(ctx, "token", token), nil
}

type TokenAuth struct {
	Token string
}

// GetRequestMetadata Return value is mapped to request headers.
func (t TokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"authentication": fmt.Sprintf("bearer %s", t.Token),
	}, nil
}

func (TokenAuth) RequireTransportSecurity() bool {
	return false
}
