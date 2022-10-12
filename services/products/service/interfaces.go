package service

import (
	"context"

	"github.com/BradErz/monorepo/services/products/models"
)

type Storage interface {
	CreateProduct(ctx context.Context, req *models.CreateProductRequest) (*models.Product, error)
	ListProducts(ctx context.Context, req *models.ListProductRequest) (*models.ListProductResponse, error)
	UpdateProduct(ctx context.Context, req *models.UpdateProductRequest) (*models.Product, error)
	GetProduct(ctx context.Context, req *models.GetProductRequest) (*models.Product, error)
}

type Cache interface {
	CreateProduct(ctx context.Context, product *models.Product) error
	GetProduct(ctx context.Context, id string) (*models.Product, error)
}
