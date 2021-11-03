package models

import "time"

type Product struct {
	Name     string `json:"name,omitempty" bson:"name"`
	ImageURL string `json:"image_url,omitempty" bson:"image_url"`

	CreateTime time.Time  `json:"create_time" bson:"create_time"`
	UpdateTime *time.Time `json:"update_time" bson:"update_time"`
	DeleteTime *time.Time `json:"delete_time" bson:"delete_time"`
}

type CreateProductRequest struct {
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
}

type ListProductRequest struct {
	Filter    string
	OrderBy   string
	PageToken string
	PageSize  int32
}

type ListProductResponse struct {
	Products      []*Product `json:"products"`
	NextPageToken string
}

type GetProductRequest struct {
	Name string
}

type UpdateProductRequest struct {
	Product *Product
	Paths   []string
}

type DeleteProductRequest struct {
	Name string
}
