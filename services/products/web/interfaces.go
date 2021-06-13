package web

import (
	"context"

	"github.com/BradErz/monorepo/services/products/models"
)

type Service interface {
	CreateProduct(ctx context.Context, req *models.CreateProductRequest) (*models.Product, error)
	GetProduct(ctx context.Context, req *models.GetProductRequest) (*models.Product, error)
	ListProducts(ctx context.Context, req *models.ListProductRequest) (*models.ListProductResponse, error)
	UpdateProduct(ctx context.Context, req *models.UpdateProductRequest) (*models.Product, error)
}
