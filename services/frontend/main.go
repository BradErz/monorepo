package main

import (
	"context"
	"fmt"

	productsv1 "github.com/BradErz/monorepo/gen/go/products/v1"

	"github.com/BradErz/monorepo/pkg/telemetry"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func main() {
	if err := app(); err != nil {
		logrus.WithError(err).Fatal("failed")
	}
}

func app() error {
	lgr := logrus.NewEntry(logrus.New())

	if err := telemetry.Init(lgr, telemetry.WithServiceName("frontend"), telemetry.WithEnabled()); err != nil {
		return fmt.Errorf("failed to setup telemetry: %w", err)
	}

	grpcOpts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	}

	productConn, err := grpc.Dial("localhost:8002", grpcOpts...)
	if err != nil {
		return fmt.Errorf("failed to dial to reviews service: %w", err)
	}
	prodClient := productsv1.NewProductsServiceClient(productConn)
	req := &productsv1.CreateProductRequest{
		Name:        "my amazing",
		ImageUrl:    "https://amazing.life/image.png",
		Description: "this really is an amazing product you would not imagine",
		Price:       9.99,
		Category:    productsv1.ProductCategory_BOOK,
	}
	resp, err := prodClient.CreateProduct(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}
	lgr.Infof("got resp %+v", resp.GetProduct())
	return nil
}

//
//func app() error {
//	lgr := logrus.NewEntry(logrus.New())
//
//	if err := telemetry.Init(lgr, telemetry.WithServiceName("frontend"), telemetry.WithEnabled()); err != nil {
//		return fmt.Errorf("failed to setup telemetry: %w", err)
//	}
//
//	grpcOpts := []grpc.DialOption{
//		grpc.WithInsecure(),
//		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
//		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
//	}
//
//	reviewConn, err := grpc.Dial("localhost:8001", grpcOpts...)
//	if err != nil {
//		return fmt.Errorf("failed to dial to reviews service: %w", err)
//	}
//	productConn, err := grpc.Dial("localhost:8002", grpcOpts...)
//	if err != nil {
//		return fmt.Errorf("failed to dial to reviews service: %w", err)
//	}
//
//	svc := &service{
//		reviewsClient:  reviewsv1.NewReviewsServiceClient(reviewConn),
//		productsClient: productsv1.NewProductsServiceClient(productConn),
//	}
//	ctx := context.Background()
//	tr := otel.Tracer("")
//	ctx, span := tr.Start(ctx, "full flow")
//	defer span.End()
//	if _, err := svc.createProduct(ctx, "ciccio"); err != nil {
//		return fmt.Errorf("failed to create product: %w", err)
//	}
//
//	if _, err := svc.createReview(ctx, "ciccio", "1"); err != nil {
//		return fmt.Errorf("failed to create review: %w", err)
//	}
//
//	resp, err := svc.getEnrichedProduct(ctx, "ciccio")
//	if err != nil {
//		return fmt.Errorf("failed to enrich product info: %w", err)
//	}
//	logrus.Infof("got enriched product: %s", resp.String())
//	return nil
//}
//
//func (svc *service) createProduct(ctx context.Context, id string) (*productsv1.Product, error) {
//	ctx, cancel := context.WithTimeout(ctx, time.Second)
//	defer cancel()
//
//	req := &productsv1.CreateProductRequest{Product: &productsv1.Product{
//		Id:       id,
//		Name:     "Ciccio Product",
//		ImageUrl: "https://imgur.com",
//	}}
//
//	return svc.productsClient.CreateProduct(ctx, req)
//}
//
//func (svc *service) createReview(ctx context.Context, productID, reviewID string) (*reviewsv1.Review, error) {
//	ctx, cancel := context.WithTimeout(ctx, time.Second)
//	defer cancel()
//
//	req := &reviewsv1.CreateReviewRequest{Review: &reviewsv1.Review{
//		Id:        reviewID,
//		ProductId: productID,
//		Title:     "my amazing review",
//		Body:      "oh my god ciccio is amazing",
//	}}
//
//	return svc.reviewsClient.CreateReview(ctx, req)
//}
//
//type service struct {
//	productsClient productsv1.ProductsServiceClient
//	reviewsClient  reviewsv1.ReviewsServiceClient
//}
//
//type EnrichedProduct struct {
//	Product *productsv1.Product
//	Reviews []*reviewsv1.Review
//}
//
//func (e EnrichedProduct) String() string {
//	var sb strings.Builder
//	sb.WriteString(fmt.Sprintf("product: %s\n", e.Product.String()))
//	for i, rev := range e.Reviews {
//		sb.WriteString(fmt.Sprintf("review %d: %s\n", i, rev.String()))
//	}
//	return sb.String()
//}
//
//func (svc *service) getEnrichedProduct(ctx context.Context, id string) (*EnrichedProduct, error) {
//	tr := otel.Tracer("")
//	ctx, span := tr.Start(ctx, "getEnrichedProduct")
//	defer span.End()
//
//	resp := &EnrichedProduct{
//		Product: &productsv1.Product{},
//		Reviews: []*reviewsv1.Review{},
//	}
//	g, ctx := errgroup.WithContext(ctx)
//
//	g.Go(func() error {
//		getProductReq := &productsv1.GetProductRequest{Id: id}
//		productResp, err := svc.productsClient.GetProduct(ctx, getProductReq)
//		if err != nil {
//			return fmt.Errorf("failed to load product: %w", err)
//		}
//		resp.Product = productResp
//		return nil
//	})
//
//	g.Go(func() error {
//		listReviewReq := &reviewsv1.ListReviewsRequest{ProductId: id}
//		reviewResp, err := svc.reviewsClient.ListReviews(ctx, listReviewReq)
//		if err != nil {
//			return fmt.Errorf("failed to load product reviews: %w", err)
//		}
//		resp.Reviews = reviewResp.Reviews
//		return nil
//	})
//
//	if err := g.Wait(); err != nil {
//		return nil, err
//	}
//
//	return resp, nil
//}
