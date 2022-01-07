package storage

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/BradErz/monorepo/services/products/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type product struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	ImageURL    string             `bson:"image_url"`
	Description string             `bson:"description"`
	Category    string             `bson:"category"`
	Price       float32            `bson:"price"`

	CreateTime time.Time  `bson:"create_time"`
	UpdateTime *time.Time `bson:"update_time"`
	DeleteTime *time.Time `bson:"delete_time"`
}

func fromCreateProductRequest(prod *models.CreateProductRequest) *product {
	return &product{
		ID:          primitive.NewObjectID(),
		Name:        prod.Name,
		ImageURL:    prod.ImageURL,
		Description: prod.Description,
		Category:    string(prod.Category),
		Price:       prod.Price,
		CreateTime:  time.Now().UTC(),
	}
}

func toProduct(prod *product) *models.Product {
	return &models.Product{
		ID:          prod.ID.Hex(),
		Name:        prod.Name,
		ImageURL:    prod.ImageURL,
		Description: prod.Description,
		Category:    models.ProductCategory(prod.Category),
		Price:       prod.Price,
		CreateTime:  prod.CreateTime,
		UpdateTime:  prod.UpdateTime,
		DeleteTime:  prod.DeleteTime,
	}
}

func toProducts(prods []*product) []*models.Product {
	res := make([]*models.Product, len(prods))
	for i, prod := range prods {
		res[i] = toProduct(prod)
	}
	return res
}

func fromListProductsRequests(req *models.ListProductRequest) bson.M {
	return bson.M{}
}

func parseUpdateProductRequest(req *models.UpdateProductRequest) (bson.M, bson.M) {
	updates := bson.M{}
	for _, path := range req.Paths {
		switch path {
		case "image_url":
			updates["image_url"] = req.Product.ImageURL
		case "name":
			updates["name"] = req.Product.Name
		case "price":
			updates["price"] = req.Product.Price
		case "category":
			updates["category"] = req.Product.Category
		}
	}
	return bson.M{"_id": req.Product.ID}, bson.M{"$set": updates}
}
