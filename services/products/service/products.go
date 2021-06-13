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

type Products struct {
	store Storage
}

func NewProducts(store Storage) *Products {
	return &Products{store: store}
}

func (p *Products) GetProduct(ctx context.Context, req *models.GetProductRequest) (*models.Product, error) {
	product, err := p.store.GetProduct(ctx, req)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Products) CreateProduct(ctx context.Context, req *models.CreateProductRequest) (*models.Product, error) {
	product, err := p.store.CreateProduct(ctx, req)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Products) ListProducts(ctx context.Context, req *models.ListProductRequest) (*models.ListProductResponse, error) {
	listResp, err := p.store.ListProducts(ctx, req)
	if err != nil {
		return nil, err
	}
	return listResp, nil
}

func (p *Products) UpdateProduct(ctx context.Context, req *models.UpdateProductRequest) (*models.Product, error) {
	product, err := p.store.UpdateProduct(ctx, req)
	if err != nil {
		return nil, err
	}
	return product, nil
}
