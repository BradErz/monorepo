package web

import (
	"context"
	"strings"

	"github.com/bufbuild/connect-go"
	"github.com/go-logr/logr"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/BradErz/monorepo/services/products/models"

	productsv1 "go.buf.build/bufbuild/connect-go/braderz/monorepo/products/v1"
	productsv1connect "go.buf.build/bufbuild/connect-go/braderz/monorepo/products/v1/productsv1connect"
)

type Server struct {
	lgr     logr.Logger
	service Service
}

var _ productsv1connect.ProductsServiceHandler = (*Server)(nil)

func New(lgr logr.Logger, service Service) *Server {
	return &Server{
		lgr:     lgr,
		service: service,
	}
}

func (s *Server) ListProducts(ctx context.Context, req *connect.Request[productsv1.ListProductsRequest]) (*connect.Response[productsv1.ListProductsResponse], error) {
	resp, err := s.service.ListProducts(ctx, toModelListProductRequest(req.Msg))
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(toProtoListProductResponse(resp)), nil
}

func (s *Server) CreateProduct(ctx context.Context, req *connect.Request[productsv1.CreateProductRequest]) (*connect.Response[productsv1.CreateProductResponse], error) {
	resp, err := s.service.CreateProduct(ctx, toModelCreateProductRequest(req.Msg))
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&productsv1.CreateProductResponse{Product: toProtoProduct(resp)}), nil
}

func (s *Server) UpdateProduct(ctx context.Context, req *connect.Request[productsv1.UpdateProductRequest]) (*connect.Response[productsv1.UpdateProductResponse], error) {
	resp, err := s.service.UpdateProduct(ctx, toModelUpdateProductRequest(req.Msg))
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&productsv1.UpdateProductResponse{Product: toProtoProduct(resp)}), nil
}

func (s *Server) GetProduct(ctx context.Context, req *connect.Request[productsv1.GetProductRequest]) (*connect.Response[productsv1.GetProductResponse], error) {
	resp, err := s.service.GetProduct(ctx, &models.GetProductRequest{ID: req.Msg.Id})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&productsv1.GetProductResponse{Product: toProtoProduct(resp)}), nil
}

// TODO: fix this :joy:
func (s *Server) GetProductOverview(ctx context.Context, req *connect.Request[productsv1.GetProductOverviewRequest]) (*connect.Response[productsv1.GetProductOverviewResponse], error) {
	return nil, nil
}

// func (srv *Server) GetProductOverview(ctx context.Context, req *productsv1.GetProductOverviewRequest) (*productsv1.GetProductOverviewResponse, error) {
// 	product, err := srv.service.GetProduct(ctx, &models.GetProductRequest{ID: req.GetProductId()})
// 	if err != nil {
// 		return nil, err
// 	}

// 	protoResp := &productsv1.GetProductOverviewResponse{
// 		ProductOverview: &productsv1.ProductOverview{
// 			Product: toProtoProduct(product),
// 		},
// 	}

// 	for _, v := range req.GetFieldMask().GetPaths() {
// 		if v == "reviews" {
// 			srv.lgr.Info("path was a reviews")
// 			listReviewReq := &reviewsv1.ListReviewsRequest{ProductId: req.GetProductId(), PageSize: 10}
// 			reviewResp, err := srv.reviewsClient.ListReviews(ctx, listReviewReq)
// 			if err != nil {
// 				return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "could not get reviews for %s", req.GetProductId())
// 			}
// 			protoResp.ProductOverview.Reviews = reviewResp.GetReviews()
// 		}
// 	}

// 	return protoResp, nil
// }

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
