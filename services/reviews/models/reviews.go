package models

import "time"

type Review struct {
	ID        string `bson:"_id"`
	ProductID string `bson:"product_id"`
	Title     string
	Body      string
	Rating    uint

	CreateTime time.Time
	UpdateTime *time.Time
	DeleteTime *time.Time
}

type CreateReviewRequest struct {
	ProductID string
	Name      string
	Title     string
	Body      string
	Rating    uint
}

type ListReviewsRequest struct {
	ProductID string
	PageSize  int32
	PageToken string
}

type ListReviewsResponse struct {
	Reviews       []*Review
	NextPageToken string
}
