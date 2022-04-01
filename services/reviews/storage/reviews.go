package storage

import (
	"context"

	"github.com/BradErz/monorepo/pkg/xerrors"
	"github.com/BradErz/monorepo/services/reviews/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Reviews struct {
	coll *mongo.Collection
}

func NewReviews(md *mongo.Database) *Reviews {
	return &Reviews{
		coll: md.Collection("reviews"),
	}
}

func (p *Reviews) CreateReview(ctx context.Context, req *models.CreateReviewRequest) (*models.Review, error) {
	review, err := fromCreateReviewRequest(req)
	if err != nil {
		return nil, err
	}

	if _, err := p.coll.InsertOne(ctx, review); err != nil {
		return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "failed to create review")
	}

	return toReview(review), nil
}

func (p *Reviews) ListReviews(ctx context.Context, req *models.ListReviewsRequest) (*models.ListReviewsResponse, error) {
	filter, opts, err := fromListReviewsRequest(req)
	if err != nil {
		return nil, err
	}

	cur, err := p.coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "failed to list products")
	}

	var reviews []*review
	if err := cur.All(ctx, &reviews); err != nil {
		return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "failed to unmarshal products")
	}

	resp := &models.ListReviewsResponse{
		Reviews: toReviews(reviews),
	}

	if len(reviews) != 0 {
		resp.NextPageToken = reviews[len(reviews)-1].ID.Hex()
	}

	return resp, nil
}
