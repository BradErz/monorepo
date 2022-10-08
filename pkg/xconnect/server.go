package xconnect

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/rs/cors"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	httpServer *http.Server
	lis        net.Listener
}

func NewServer(handler http.Handler, lis net.Listener) *Server {
	c := cors.Default()
	return &Server{
		lis: lis,
		httpServer: &http.Server{
			Handler: h2c.NewHandler(
				otelhttp.NewHandler(
					c.Handler(handler),
					"http.server",
					otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
						return operation + " " + r.URL.Path
					}),
				),
				&http2.Server{},
			),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	if err := s.httpServer.Serve(s.lis); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Addr() string {
	return s.lis.Addr().String()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
