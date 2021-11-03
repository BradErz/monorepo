package storage

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/BradErz/monorepo/pkg/xerrors"

	"github.com/BradErz/monorepo/services/products/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Products struct {
	coll *mongo.Collection
}

func NewProducts(md *mongo.Database) (*Products, error) {
	return &Products{
		coll: md.Collection("products"),
	}, nil
}

func (p *Products) GetProduct(ctx context.Context, req *models.GetProductRequest) (*models.Product, error) {
	filter := bson.M{"name": req.Name}
	product := &models.Product{}
	if err := p.coll.FindOne(ctx, filter).Decode(&product); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, xerrors.Wrapf(xerrors.CodeNotFound, err, "product with name %s does not exist", req.Name)
		}
		return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "failed to find product")
	}
	return product, nil
}

func (p *Products) CreateProduct(ctx context.Context, req *models.CreateProductRequest) (*models.Product, error) {
	product := newProductFromCreate(req)
	if _, err := p.coll.InsertOne(ctx, product); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, xerrors.Wrapf(xerrors.CodeAlreadyExists, err, "name %s already exists", req.Name)
		}
		return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "failed to create product")
	}

	return product, nil
}

func (p *Products) ListProducts(ctx context.Context, req *models.ListProductRequest) (*models.ListProductResponse, error) {
	cur, err := p.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "failed to list products")
	}
	var products []*models.Product
	if err := cur.All(ctx, &products); err != nil {
		return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "failed to unmarshal products")
	}
	return &models.ListProductResponse{
		Products:      products,
		NextPageToken: "111",
	}, nil
}

func (p *Products) UpdateProduct(ctx context.Context, req *models.UpdateProductRequest) (*models.Product, error) {
	opts := options.FindOneAndUpdate()
	opts.SetReturnDocument(options.After)

	filter, updates := parseUpdateProductRequest(req)
	product := &models.Product{}
	if err := p.coll.FindOneAndUpdate(ctx, filter, updates, opts).Decode(&product); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, xerrors.Wrapf(xerrors.CodeNotFound, err, "no product found with name: %s", req.Product.Name)
		}
		return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "failed to update product: %s", req.Product.Name)
	}

	return product, nil
}

func parseUpdateProductRequest(req *models.UpdateProductRequest) (bson.M, bson.M) {
	updates := bson.M{}
	for _, path := range req.Paths {
		switch path {

		case "image_url":
			updates["image_url"] = req.Product.ImageURL
		}
	}
	return bson.M{"name": req.Product.Name}, bson.M{"$set": updates}
}

func newProductFromCreate(req *models.CreateProductRequest) *models.Product {
	now := time.Now().UTC()
	return &models.Product{
		Name:       req.Name,
		ImageURL:   req.ImageURL,
		CreateTime: now,
	}
}
