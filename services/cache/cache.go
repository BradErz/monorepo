package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/BradErz/monorepo/pkg/xcache"
	"github.com/BradErz/monorepo/pkg/xerrors"
	"github.com/BradErz/monorepo/services/products/models"
	"github.com/go-redis/redis/v9"
)

type Products struct {
	client *xcache.Client
}

func NewProducts(client *xcache.Client) *Products {
	return &Products{
		client: client,
	}
}

func (p *Products) CreateProduct(ctx context.Context, product *models.Product) error {
	b, err := json.Marshal(product)
	if err != nil {
		return xerrors.Wrapf(xerrors.CodeInvalidArgument, err, "failed to save %s to cache", product.ID)
	}
	if _, err := p.client.Client.Set(ctx, p.getKeyByID(product.ID), b, time.Minute).Result(); err != nil {
		return xerrors.Wrapf(xerrors.CodeUnavailable, err, "failed to save %s to cache", product.ID)
	}
	return nil
}

func (p *Products) GetProduct(ctx context.Context, id string) (*models.Product, error) {
	res, err := p.client.Client.Get(ctx, p.getKeyByID(id)).Bytes()
	switch {
	case err == redis.Nil:
		return nil, xerrors.Wrapf(xerrors.CodeNotFound, err, "could not find product %s in the cache", id)
	case err != nil:
		return nil, xerrors.Wrapf(xerrors.CodeUnavailable, err, "failed to save query redis for product %s", id)
	}
	product := &models.Product{}
	if err := json.Unmarshal(res, &product); err != nil {
		return nil, xerrors.Wrapf(xerrors.CodeInvalidArgument, err, "failed to unmarshal product id %s from redis", id)
	}
	return product, nil
}

func (p *Products) getKeyByID(id string) string {
	return p.client.Prefix("product:" + id)
}
