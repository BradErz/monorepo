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

func getAllCodes() map[xerrors.Code]codes.Code {
	return map[xerrors.Code]codes.Code{
		xerrors.CodeUnknown:         codes.Unknown,
		xerrors.CodeNotFound:        codes.NotFound,
		xerrors.CodeInvalidArgument: codes.InvalidArgument,
		xerrors.CodeAlreadyExists:   codes.AlreadyExists,
		xerrors.CodeUnauthenticated: codes.Unauthenticated,
		xerrors.CodeInternal:        codes.Internal,
	}
}

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

	c, ok := getAllCodes()[myError.Code()]
	if !ok {
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
