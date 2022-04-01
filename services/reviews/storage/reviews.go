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

	// If we have no reviews from the database we have no next page token
	// AND
	// if the amount of the reviews we get back is equal to the page size then we might have more reviews.
	//
	// maybe instead we should always do a +1 on the page size and return that ID as the actual "next" page token
	// because then if someone requests 1 page of 10 docs, but we only have 10 docs currently we will still give a
	// next page token of the 10th document due to the "$gt" in the filter which means they will query again but get an
	// empty page.
	if len(reviews) != 0 && len(reviews) == int(req.PageSize) {
		resp.NextPageToken = reviews[len(reviews)-1].ID.Hex()
	}

	return resp, nil
}
