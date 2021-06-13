package service

import (
	"context"

	"github.com/BradErz/monorepo/services/reviews/models"
)

type Storage interface {
	CreateReview(ctx context.Context, req *models.CreateReviewRequest) (*models.Review, error)
	ListReviews(ctx context.Context, req *models.ListReviewsRequest) (*models.ListReviewsResponse, error)
}

type Reviews struct {
	store Storage
}

func NewReviews(store Storage) *Reviews {
	return &Reviews{store: store}
}

func (r *Reviews) CreateReview(ctx context.Context, req *models.CreateReviewRequest) (*models.Review, error) {
	product, err := r.store.CreateReview(ctx, req)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *Reviews) ListReviews(ctx context.Context, req *models.ListReviewsRequest) (*models.ListReviewsResponse, error) {
	listResp, err := r.store.ListReviews(ctx, req)
	if err != nil {
		return nil, err
	}
	return listResp, nil
}
