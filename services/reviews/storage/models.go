package storage

import (
	"time"

	"github.com/BradErz/monorepo/pkg/xerrors"
	"github.com/BradErz/monorepo/services/reviews/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type review struct {
	ID         primitive.ObjectID `bson:"_id"`
	ProductID  primitive.ObjectID `bson:"product_id"`
	Title      string             `bson:"title"`
	Body       string             `bson:"body"`
	Rating     uint               `bson:"rating"`
	CreateTime time.Time          `bson:"create_time"`
	UpdateTime *time.Time         `bson:"update_time"`
	DeleteTime *time.Time         `bson:"delete_time"`
}

func fromCreateReviewRequest(req *models.CreateReviewRequest) (*review, error) {
	prodID, err := primitive.ObjectIDFromHex(req.ProductID)
	if err != nil {
		return nil, xerrors.NotValidObjectID(req.ProductID)
	}
	return &review{
		ID:         primitive.NewObjectID(),
		ProductID:  prodID,
		Title:      req.Title,
		Body:       req.Body,
		Rating:     req.Rating,
		CreateTime: time.Now().UTC(),
	}, nil
}

func toReview(rev *review) *models.Review {
	return &models.Review{
		ID:         rev.ID.Hex(),
		ProductID:  rev.ProductID.Hex(),
		Title:      rev.Title,
		Body:       rev.Body,
		Rating:     rev.Rating,
		CreateTime: rev.CreateTime,
		UpdateTime: rev.UpdateTime,
		DeleteTime: rev.DeleteTime,
	}
}

func toReviews(revs []*review) []*models.Review {
	res := make([]*models.Review, len(revs))
	for i, rev := range revs {
		res[i] = toReview(rev)
	}
	return res
}

func fromListReviewsRequest(req *models.ListReviewsRequest) (bson.M, *options.FindOptions, error) {
	prodID, err := primitive.ObjectIDFromHex(req.ProductID)
	if err != nil {
		return nil, nil, xerrors.NotValidObjectID(req.ProductID)
	}

	filter := bson.M{
		"product_id": prodID,
	}

	if req.PageToken != "" {
		pageToken, err := primitive.ObjectIDFromHex(req.PageToken)
		if err != nil {
			return nil, nil, xerrors.NotValidObjectID(req.PageToken)
		}

		filter["_id"] = bson.M{"$gt": pageToken}
	}

	opts := options.Find().
		SetLimit(int64(req.PageSize))

	return filter, opts, nil
}
