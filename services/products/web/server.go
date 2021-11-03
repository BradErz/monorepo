package web

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"github.com/BradErz/monorepo/pkg/xerrors"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/BradErz/monorepo/services/products/models"

	productsv1 "github.com/BradErz/monorepo/gen/go/org/products/v1"

	"github.com/BradErz/monorepo/pkg/xgrpc"
	"github.com/sirupsen/logrus"
)

type Server struct {
	le      *logrus.Entry
	service Service
}

var _ productsv1.ProductsServiceServer = (*Server)(nil)

func New(le *logrus.Entry, service Service) (*Server, error) {
	return &Server{
		le:      le,
		service: service,
	}, nil
}

func Register(productsSrv productsv1.ProductsServiceServer) xgrpc.RegisterServerFunc {
	return func(s *grpc.Server) {
		productsv1.RegisterProductsServiceServer(s, productsSrv)
	}
}

func (srv *Server) GetProduct(ctx context.Context, req *productsv1.GetProductRequest) (*productsv1.Product, error) {
	product, err := srv.service.GetProduct(ctx, &models.GetProductRequest{Name: req.GetId()})
	if err != nil {
		return nil, err
	}
	return toProtoProduct(product), nil
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

func (srv *Server) CreateProduct(ctx context.Context, req *productsv1.CreateProductRequest) (*productsv1.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	product, err := srv.service.CreateProduct(ctx, toModelCreateProductRequest(req))
	if err != nil {
		return nil, err
	}
	return toProtoProduct(product), nil
}

func (srv *Server) UpdateProduct(ctx context.Context, req *productsv1.UpdateProductRequest) (*productsv1.Product, error) {
	if !req.GetFieldMask().IsValid(req.GetProduct()) {
		return nil, xerrors.Newf(xerrors.CodeInvalidArgument, "specified paths in field_mask are invalid")
	}
	product, err := srv.service.UpdateProduct(ctx, toModelUpdateProductRequest(req))
	if err != nil {
		return nil, err
	}

	return toProtoProduct(product), nil
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
		Name:     req.GetProduct().GetName(),
		ImageURL: req.GetProduct().GetImageUrl(),
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
		Name:       product.Name,
		CreateTime: timestamppb.New(product.CreateTime),
		ImageUrl:   product.ImageURL,
	}
	if product.UpdateTime != nil {
		p.UpdateTime = timestamppb.New(*product.UpdateTime)
	}
	if product.DeleteTime != nil {
		p.UpdateTime = timestamppb.New(*product.DeleteTime)
	}
	return p
}
