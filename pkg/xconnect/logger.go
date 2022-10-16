package xconnect

import (
	"context"
	"errors"
	"time"

	"github.com/BradErz/monorepo/pkg/xerrors"
	"github.com/bufbuild/connect-go"
	"github.com/go-logr/logr"
)

type httpLog struct {
	Duration      time.Duration `json:"duration"`
	Path          string        `json:"path"`
	XForwardedFor string        `json:"x-forwarded-for"`
}

func LogrInterceptor(lgr logr.Logger) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			start := time.Now()

			resp, err := next(ctx, req)

			fields := []any{
				"http", &httpLog{
					Duration:      time.Since(start),
					Path:          req.Spec().Procedure,
					XForwardedFor: req.Header().Get("X-Forwarded-For"),
				},
				"trace_id", req.Header().Get("X-B3-Traceid"),
				"span_id", req.Header().Get("X-B3-Spanid"),
				"request_id", req.Header().Get("x-request-id"),
			}

			if err != nil {
				if myError := new(xerrors.Error); errors.As(err, &myError) {
					fields = append(fields,
						"error_details", myError,
					)
				}
				lgr.Error(err, "request had error", fields...)
				return resp, err
			}

			lgr.Info("request completed", fields...)
			return resp, err
		}
	}

	return connect.UnaryInterceptorFunc(interceptor)
}
