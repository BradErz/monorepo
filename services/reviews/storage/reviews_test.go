package storage

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/BradErz/monorepo/pkg/xmongo"
	"github.com/BradErz/monorepo/services/reviews/models"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestListReviews(t *testing.T) {
	t.Run("amazing life", func(t *testing.T) {
		md, clear := xmongo.TestingSetup(t, "my-test")
		defer clear()
		store := NewReviews(md.Database)
		revs := getReviews(9)
		for _, rev := range revs {
			_, err := store.CreateReview(context.Background(), rev)
			require.NoError(t, err)
		}

		nextPageToken := ""
		for {
			req := &models.ListReviewsRequest{ProductID: revs[0].ProductID, PageSize: 2}
			if nextPageToken != "" {
				req.PageToken = nextPageToken
			}

			resp, err := store.ListReviews(context.Background(), req)
			require.NoError(t, err)
			for i, v := range resp.Reviews {
				t.Logf("%d: %+v", i, v)
			}

			nextPageToken = resp.NextPageToken
			if resp.NextPageToken == "" {
				break
			}
		}
	})
}

func getReviews(count int) []*models.CreateReviewRequest {
	prodID := primitive.NewObjectID().Hex()
	res := make([]*models.CreateReviewRequest, count)
	for i := 0; i < count; i++ {
		res[i] = &models.CreateReviewRequest{
			ProductID: prodID,
			Name:      fmt.Sprintf("review name: %d", i),
			Title:     fmt.Sprintf("review title: %d", i),
			Body:      fmt.Sprintf("review body: %d", i),
			Rating:    uint(rand.Intn(5)),
		}
	}
	return res
}
