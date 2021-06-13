package xerrors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestError(t *testing.T) {
	t.Run("Newf", func(t *testing.T) {
		err := Newf(CodeInvalidArgument, "name: %s is invalid", "ciccio")

		myError := new(Error)
		require.True(t, errors.As(err, &myError))

		require.Equal(t, CodeInvalidArgument, myError.Code())
		require.Nil(t, myError.Details())
	})
	t.Run("Wrapf", func(t *testing.T) {
		wrappedErr := errors.New("my wrapped error")
		err := Wrapf(CodeInvalidArgument, wrappedErr, "name: %s is invalid", "ciccio")

		myError := new(Error)
		require.True(t, errors.As(err, &myError))

		require.Equal(t, CodeInvalidArgument, myError.Code())
		require.Nil(t, myError.Details())
		require.True(t, errors.Is(err, wrappedErr))
	})

	t.Run("Details", func(t *testing.T) {
		wrappedErr := errors.New("my wrapped error")
		details := map[string]string{
			"name_limit": "100",
		}
		err := Detailsf(CodeInvalidArgument, wrappedErr, details, "name: %s is invalid", "ciccio")

		myError := new(Error)
		require.True(t, errors.As(err, &myError))

		require.Equal(t, CodeInvalidArgument, myError.Code())
		require.True(t, errors.Is(err, wrappedErr))
		require.Equal(t, details, myError.Details())
	})
}
