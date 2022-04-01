package web

import (
	"context"
	"strings"
	"time"

	"github.com/go-logr/logr"

	"google.golang.org/grpc"

	"github.com/BradErz/monorepo/pkg/xerrors"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/BradErz/monorepo/services/products/models"

	productsv1 "github.com/BradErz/monorepo/gen/go/products/v1"
	reviewsv1 "github.com/BradErz/monorepo/gen/go/reviews/v1"

	"github.com/BradErz/monorepo/pkg/xgrpc"
)

type Server struct {
	lgr           logr.Logger
	service       Service
	reviewsClient reviewsv1.ReviewsServiceClient
}

var _ productsv1.ProductsServiceServer = (*Server)(nil)

func New(lgr logr.Logger, service Service, reviewsClient reviewsv1.ReviewsServiceClient) (*Server, error) {
	return &Server{
		lgr:           lgr,
		service:       service,
		reviewsClient: reviewsClient,
	}, nil
}

func Register(productsSrv productsv1.ProductsServiceServer) xgrpc.RegisterServerFunc {
	return func(s *grpc.Server) {
		productsv1.RegisterProductsServiceServer(s, productsSrv)
	}
}

func (srv *Server) GetProduct(ctx context.Context, req *productsv1.GetProductRequest) (*productsv1.GetProductResponse, error) {
	product, err := srv.service.GetProduct(ctx, &models.GetProductRequest{ID: req.GetId()})
	if err != nil {
		return nil, err
	}
	return &productsv1.GetProductResponse{Product: toProtoProduct(product)}, nil
}

func (srv *Server) GetProductOverview(ctx context.Context, req *productsv1.GetProductOverviewRequest) (*productsv1.GetProductOverviewResponse, error) {
	product, err := srv.service.GetProduct(ctx, &models.GetProductRequest{ID: req.GetProductId()})
	if err != nil {
		return nil, err
	}

	protoResp := &productsv1.GetProductOverviewResponse{
		ProductOverview: &productsv1.ProductOverview{
			Product: toProtoProduct(product),
		},
	}

	for _, v := range req.GetFieldMask().GetPaths() {
		if v == "reviews" {
			srv.lgr.Info("path was a reviews")
			listReviewReq := &reviewsv1.ListReviewsRequest{ProductId: req.GetProductId(), PageSize: 10}
			reviewResp, err := srv.reviewsClient.ListReviews(ctx, listReviewReq)
			if err != nil {
				return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "could not get reviews for %s", req.GetProductId())
			}
			protoResp.ProductOverview.Reviews = reviewResp.GetReviews()
		}
	}

	return protoResp, nil
}

func (srv *Server) ListProducts(ctx context.Context, req *productsv1.ListProductsRequest) (*productsv1.ListProductsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	listResp, err := srv.service.ListProducts(ctx, toModelListProductRequest(req))
	if err != nil {
		return nil, err
	}

	return toProtoListProductResponse(listResp), nil
}

func (srv *Server) CreateProduct(ctx context.Context, req *productsv1.CreateProductRequest) (*productsv1.CreateProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	if err := req.Validate(); err != nil {
		return nil, xerrors.Newf(xerrors.CodeInvalidArgument, err.Error())
	}

	product, err := srv.service.CreateProduct(ctx, toModelCreateProductRequest(req))
	if err != nil {
		return nil, err
	}
	return &productsv1.CreateProductResponse{Product: toProtoProduct(product)}, nil
}

func (srv *Server) UpdateProduct(ctx context.Context, req *productsv1.UpdateProductRequest) (*productsv1.UpdateProductResponse, error) {
	if !req.GetFieldMask().IsValid(req.GetProduct()) {
		return nil, xerrors.Newf(xerrors.CodeInvalidArgument, "specified paths in field_mask are invalid")
	}
	product, err := srv.service.UpdateProduct(ctx, toModelUpdateProductRequest(req))
	if err != nil {
		return nil, err
	}

	return &productsv1.UpdateProductResponse{Product: toProtoProduct(product)}, nil
}

func toModelUpdateProductRequest(req *productsv1.UpdateProductRequest) *models.UpdateProductRequest {
	return &models.UpdateProductRequest{
		Product: &models.Product{
			Name:     req.GetProduct().GetName(),
			ImageURL: req.GetProduct().GetImageUrl(),
		},
		Paths: req.FieldMask.GetPaths(),
	}
}

func toModelCreateProductRequest(req *productsv1.CreateProductRequest) *models.CreateProductRequest {
	return &models.CreateProductRequest{
		Name:        req.GetName(),
		ImageURL:    req.GetImageUrl(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
		Category:    models.ProductCategory(strings.ToLower(req.GetCategory().String())),
	}
}

func toModelListProductRequest(req *productsv1.ListProductsRequest) *models.ListProductRequest {
	return &models.ListProductRequest{
		Filter:    req.GetFilter(),
		OrderBy:   req.GetOrderBy(),
		PageToken: req.GetPageToken(),
		PageSize:  req.GetPageSize(),
	}
}

func toProtoListProductResponse(productResp *models.ListProductResponse) *productsv1.ListProductsResponse {
	resp := make([]*productsv1.Product, len(productResp.Products))
	for i, product := range productResp.Products {
		resp[i] = toProtoProduct(product)
	}
	return &productsv1.ListProductsResponse{
		Products:      resp,
		NextPageToken: productResp.NextPageToken,
	}
}

func toProtoProduct(product *models.Product) *productsv1.Product {
	p := &productsv1.Product{
		Id:          product.ID,
		CreateTime:  timestamppb.New(product.CreateTime),
		Name:        product.Name,
		ImageUrl:    product.ImageURL,
		Description: product.Description,
		Price:       product.Price,
		Category:    productsv1.ProductCategory(productsv1.ProductCategory_value[strings.ToUpper(string(product.Category))]),
	}
	if product.UpdateTime != nil {
		p.UpdateTime = timestamppb.New(*product.UpdateTime)
	}
	if product.DeleteTime != nil {
		p.UpdateTime = timestamppb.New(*product.DeleteTime)
	}
	return p
}
