package storage

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/BradErz/monorepo/pkg/xerrors"

	"github.com/BradErz/monorepo/services/products/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Products struct {
	coll *mongo.Collection
}

func NewProducts(md *mongo.Database) *Products {
	return &Products{
		coll: md.Collection("products"),
	}
}

func (p *Products) GetProduct(ctx context.Context, req *models.GetProductRequest) (*models.Product, error) {
	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return nil, xerrors.NotValidObjectID(req.ID)
	}
	filter := bson.M{"_id": id}
	product := &product{}
	if err := p.coll.FindOne(ctx, filter).Decode(&product); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, xerrors.Wrapf(xerrors.CodeNotFound, err, "product %s does not exist", req.ID)
		}
		return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "failed to find product")
	}
	return toProduct(product), nil
}

func (p *Products) CreateProduct(ctx context.Context, req *models.CreateProductRequest) (*models.Product, error) {
	prod := fromCreateProductRequest(req)
	if _, err := p.coll.InsertOne(ctx, prod); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, xerrors.Wrapf(xerrors.CodeAlreadyExists, err, "name %s already exists", req.Name)
		}
		return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "failed to create product")
	}

	return toProduct(prod), nil
}

func (p *Products) ListProducts(ctx context.Context, req *models.ListProductRequest) (*models.ListProductResponse, error) {
	filter := fromListProductsRequests(req)
	cur, err := p.coll.Find(ctx, filter)
	if err != nil {
		return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "failed to list products")
	}
	var prods []*product
	if err := cur.All(ctx, &prods); err != nil {
		return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "failed to unmarshal products: %v", req)
	}
	return &models.ListProductResponse{
		Products:      toProducts(prods),
		NextPageToken: "111",
	}, nil
}

func (p *Products) UpdateProduct(ctx context.Context, req *models.UpdateProductRequest) (*models.Product, error) {
	opts := options.FindOneAndUpdate()
	opts.SetReturnDocument(options.After)

	filter, updates := parseUpdateProductRequest(req)
	prod := &product{}
	if err := p.coll.FindOneAndUpdate(ctx, filter, updates, opts).Decode(&prod); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, xerrors.Wrapf(xerrors.CodeNotFound, err, "no product found with name: %s", req.Product.Name)
		}
		return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "failed to update product: %s", req.Product.Name)
	}

	return toProduct(prod), nil
}
