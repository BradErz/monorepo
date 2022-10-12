package xcache

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	client, err := New()
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.Client.Ping(ctx).Result()
	require.NoError(t, err)
	t.Log(res)
}
