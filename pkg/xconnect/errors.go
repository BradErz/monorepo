package xconnect

import (
	"context"
	"errors"

	"github.com/BradErz/monorepo/pkg/xerrors"
	"github.com/bufbuild/connect-go"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func ErrorsInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			resp, err := next(ctx, req)
			return resp, mapError(err)
		}
	}

	return connect.UnaryInterceptorFunc(interceptor)
}

func mapError(err error) error {
	if err == nil {
		return nil
	}

	myError := new(xerrors.Error)
	if !errors.As(err, &myError) {
		return err
	}

	c, ok := getAllCodes()[myError.Code]
	if !ok {
		c = connect.CodeUnknown
	}

	// we dont want to expose implementation details which might be in the myError.error to
	// the api so we always return the message
	connectError := connect.NewError(c, errors.New(myError.Message))
	if myError.Details != nil {
		detail, _ := connect.NewErrorDetail(&errdetails.ErrorInfo{
			Reason:   c.String(),
			Metadata: myError.Details,
		})
		connectError.AddDetail(detail)
	}
	return connectError
}

func getAllCodes() map[xerrors.Code]connect.Code {
	return map[xerrors.Code]connect.Code{
		xerrors.CodeUnknown:         connect.CodeUnknown,
		xerrors.CodeNotFound:        connect.CodeNotFound,
		xerrors.CodeInvalidArgument: connect.CodeInvalidArgument,
		xerrors.CodeAlreadyExists:   connect.CodeAlreadyExists,
		xerrors.CodeUnauthenticated: connect.CodeUnauthenticated,
		xerrors.CodeInternal:        connect.CodeInternal,
	}
}
