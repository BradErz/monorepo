package service

import (
	"context"

	"github.com/BradErz/monorepo/pkg/xerrors"
	"github.com/BradErz/monorepo/services/products/models"
	"github.com/go-logr/logr"
)

type Products struct {
	store Storage
	cache Cache
	lgr   logr.Logger
}

func NewProducts(lgr logr.Logger, store Storage, cache Cache) *Products {
	return &Products{
		lgr:   lgr.WithName("service.Products"),
		store: store,
		cache: cache,
	}
}

func (p *Products) GetProduct(ctx context.Context, req *models.GetProductRequest) (*models.Product, error) {
	product, err := p.cache.GetProduct(ctx, req.ID)
	switch {
	case xerrors.IsNotFound(err):
		break
	case err == nil:
		return product, nil
	}

	product, err = p.store.GetProduct(ctx, req)
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

	go func(ctx context.Context, product *models.Product) {
		if err := p.cache.CreateProduct(ctx, product); err != nil {
			p.lgr.Error(err, "can't store product in cache")
		}
	}(ctx, product)

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
