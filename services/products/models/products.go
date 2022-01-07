package models

import "time"

type ProductCategory string

const (
	ProductUnknown  ProductCategory = "product_unknown"
	ProductSoftware ProductCategory = "software"
	ProductBook     ProductCategory = "book"
	ProductCourse   ProductCategory = "course"
)

type Price struct {
}

type Product struct {
	ID          string
	Name        string
	ImageURL    string
	Description string
	Category    ProductCategory
	Price       float32

	CreateTime time.Time
	UpdateTime *time.Time
	DeleteTime *time.Time
}

type CreateProductRequest struct {
	Name        string
	ImageURL    string
	Description string
	Price       float32
	Category    ProductCategory
}

type ListProductRequest struct {
	Filter    string
	OrderBy   string
	PageToken string
	PageSize  int32
}

type ListProductResponse struct {
	Products      []*Product
	NextPageToken string
}

type GetProductRequest struct {
	ID string
}

type UpdateProductRequest struct {
	Product *Product
	Paths   []string
}

type DeleteProductRequest struct {
	Name string
}
