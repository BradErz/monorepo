package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/BradErz/monorepo/pkg/xlogger"

	"google.golang.org/protobuf/types/known/fieldmaskpb"

	productsv1 "github.com/BradErz/monorepo/gen/go/products/v1"
	reviewsv1 "github.com/BradErz/monorepo/gen/go/reviews/v1"

	"github.com/BradErz/monorepo/pkg/telemetry"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func main() {
	if err := app(); err != nil {
		log.Fatalf("failed to start application: %s", err)
	}
}

func app() error {
	lgr, err := xlogger.New()
	if err != nil {
		return fmt.Errorf("failed to create xlogger: %w", err)
	}

	tracer, err := telemetry.Init(lgr, telemetry.WithServiceName("frontend"), telemetry.WithEnabled())
	if err != nil {
		return fmt.Errorf("failed to setup telemetry: %w", err)
	}

	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	}

	productConn, err := grpc.Dial("localhost:8002", grpcOpts...)
	if err != nil {
		return fmt.Errorf("failed to dial to products service: %w", err)
	}
	reviewsConn, err := grpc.Dial("localhost:8001", grpcOpts...)
	if err != nil {
		return fmt.Errorf("failed to dial to reviews service: %w", err)
	}

	productsClient := productsv1.NewProductsServiceClient(productConn)
	reviewsClient := reviewsv1.NewReviewsServiceClient(reviewsConn)

	ctx := context.Background()

	ctx, span := tracer.Start(ctx, "hello-span")
	defer span.End()

	req := &productsv1.CreateProductRequest{
		Name:        "my amazing",
		ImageUrl:    "https://amazing.life/image.png",
		Description: "this really is an amazing product you would not imagine",
		Price:       9.99,
		Category:    productsv1.ProductCategory_PRODUCT_CATEGORY_BOOK,
	}
	resp, err := productsClient.CreateProduct(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}
	lgr.Info("got resp from creating product", "resp", resp.GetProduct())

	createReviewReq := &reviewsv1.CreateReviewRequest{
		ProductId: resp.GetProduct().GetId(),
		Title:     "this is my amazing review",
		Body:      "this product really is a life saver i do not know what i would do without it",
		Rating:    5,
	}

	createReviewResp, err := reviewsClient.CreateReview(ctx, createReviewReq)
	if err != nil {
		return fmt.Errorf("failed to create review for %s: %w", createReviewReq.GetProductId(), err)
	}

	lgr.Info("got resp from creating review", "resp", createReviewResp.GetReview())

	overviewReq := &productsv1.GetProductOverviewRequest{ProductId: createReviewReq.GetProductId(), FieldMask: &fieldmaskpb.FieldMask{Paths: []string{"reviews"}}}
	overviewResp, err := productsClient.GetProductOverview(ctx, overviewReq)
	if err != nil {
		return fmt.Errorf("failed to create review for %s: %w", createReviewReq.GetProductId(), err)
	}

	lgr.Info("got resp from overview review", "resp", overviewResp.GetProductOverview())

	return nil
}
