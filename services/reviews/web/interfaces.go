package web

import (
	"context"

	"github.com/BradErz/monorepo/services/reviews/models"
)

type Service interface {
	CreateReview(ctx context.Context, req *models.CreateReviewRequest) (*models.Review, error)
	ListReviews(ctx context.Context, req *models.ListReviewsRequest) (*models.ListReviewsResponse, error)
}
