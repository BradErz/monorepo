package models

import "time"

type Review struct {
	Name  string `json:"name,omitempty" bson:"name"`
	Title string `json:"title,omitempty" bson:"title"`
	Body  string `json:"body,omitempty" bson:"body"`

	CreateTime time.Time  `json:"create_time" bson:"create_time"`
	UpdateTime *time.Time `json:"update_time" bson:"update_time"`
	DeleteTime *time.Time `json:"delete_time" bson:"delete_time"`
}

type CreateReviewRequest struct {
	Name  string
	Title string
	Body  string
}

type ListReviewsRequest struct {
	Parent    string
	PageSize  int32
	PageToken string
}

type ListReviewsResponse struct {
	Reviews       []*Review
	NextPageToken string
}
