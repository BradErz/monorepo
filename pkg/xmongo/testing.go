package xmongo

import (
	"context"
	"testing"
	"time"

	"github.com/BradErz/monorepo/pkg/xlogger"

	"github.com/stretchr/testify/require"
)

func TestingSetup(t testing.TB, dbName string, opts ...Option) (svc *Service, tearDown func()) {
	t.Helper()

	lgr, err := xlogger.New()
	require.NoError(t, err)

	svc, err = New(lgr, dbName, opts...)
	require.NoError(t, err)

	return svc, func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		require.NoError(t, svc.Database.Drop(ctx))
	}
}
