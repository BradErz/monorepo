package storage

import (
	"context"
	"time"

	"github.com/BradErz/monorepo/pkg/xerrors"
	"github.com/BradErz/monorepo/services/reviews/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Reviews struct {
	coll *mongo.Collection
}

func NewReviews(md *mongo.Database) (*Reviews, error) {
	return &Reviews{
		coll: md.Collection("reviews"),
	}, nil
}

func (p *Reviews) CreateReview(ctx context.Context, req *models.CreateReviewRequest) (*models.Review, error) {
	product := newReviewFromCreate(req)
	if _, err := p.coll.InsertOne(ctx, product); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, xerrors.Wrapf(xerrors.CodeAlreadyExists, err, "name %s already exists", req.Name)
		}
		return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "failed to create product")
	}

	return product, nil
}

func (p *Reviews) ListReviews(ctx context.Context, req *models.ListReviewsRequest) (*models.ListReviewsResponse, error) {
	cur, err := p.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "failed to list products")
	}
	var products []*models.Review
	if err := cur.All(ctx, &products); err != nil {
		return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "failed to unmarshal products")
	}
	return &models.ListReviewsResponse{
		Reviews:       products,
		NextPageToken: "111",
	}, nil
}

func newReviewFromCreate(req *models.CreateReviewRequest) *models.Review {
	now := time.Now().UTC()
	return &models.Review{
		Name:       req.Name,
		CreateTime: now,
		Title:      req.Title,
		Body:       req.Body,
	}
}
