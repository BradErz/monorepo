package xgrpc

import (
	"context"
	"errors"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"github.com/BradErz/monorepo/pkg/xerrors"

	"google.golang.org/grpc"
)

func ErrorMapping(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	return resp, mapCodes(err)
}

func mapCodes(err error) error {
	if err == nil {
		return nil
	}

	myError := new(xerrors.Error)
	if !errors.As(err, &myError) {
		return err
	}

	var c codes.Code
	switch myError.Code() {
	case xerrors.CodeNotFound:
		c = codes.NotFound
	case xerrors.CodeInvalidArgument:
		c = codes.InvalidArgument
	case xerrors.CodeAlreadyExists:
		c = codes.AlreadyExists
	case xerrors.CodeUnauthenticated:
		c = codes.Unauthenticated
	case xerrors.CodeInternal:
		c = codes.Internal
	default:
		c = codes.Unknown
	}
	st := status.New(c, myError.Msg())
	if myError.Details() != nil {
		st, _ = st.WithDetails(&errdetails.ErrorInfo{
			Reason:   c.String(),
			Domain:   "products",
			Metadata: myError.Details(),
		})
	}

	return st.Err()
}
