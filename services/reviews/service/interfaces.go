package service

import (
	"context"

	"github.com/BradErz/monorepo/services/reviews/models"
)

//go:generate mockgen -source interfaces.go -destination mocks/storage_mock.go
type Storage interface {
	CreateReview(ctx context.Context, req *models.CreateReviewRequest) (*models.Review, error)
	ListReviews(ctx context.Context, req *models.ListReviewsRequest) (*models.ListReviewsResponse, error)
}
