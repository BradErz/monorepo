package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/BradErz/monorepo/pkg/xlogger"
// 	"github.com/bufbuild/connect-go"
// 	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

// 	productsv1 "github.com/BradErz/monorepo/gen/go/products/v1"
// 	"github.com/BradErz/monorepo/gen/go/products/v1/productsv1connect"
// 	reviewsv1 "github.com/BradErz/monorepo/gen/go/reviews/v1"
// 	"github.com/BradErz/monorepo/gen/go/reviews/v1/reviewsv1connect"

// 	"github.com/BradErz/monorepo/pkg/telemetry"
// )

// func main() {
// 	if err := app(); err != nil {
// 		log.Fatalf("failed to start application: %s", err)
// 	}
// }

// func app() error {
// 	lgr, err := xlogger.New()
// 	if err != nil {
// 		return fmt.Errorf("failed to create xlogger: %w", err)
// 	}

// 	tracer, stop, err := telemetry.Init(lgr, telemetry.WithServiceName("frontend"), telemetry.WithEnabled())
// 	if err != nil {
// 		return fmt.Errorf("failed to setup telemetry: %w", err)
// 	}
// 	defer stop(context.Background())

// 	client := &http.Client{
// 		Timeout: time.Second,
// 		Transport: otelhttp.NewTransport(http.DefaultTransport,
// 			// operation is always sent as an empty sting...
// 			otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
// 				return "http.client " + r.URL.Path
// 			}),
// 		),
// 	}

// 	productsClient := productsv1connect.NewProductsServiceClient(
// 		client,
// 		"http://localhost:8002",
// 	)

// 	reviewsClient := reviewsv1connect.NewReviewsServiceClient(
// 		client,
// 		"http://localhost:8001",
// 	)

// 	ctx := context.Background()

// 	ctx, span := tracer.Start(ctx, "hello-span")
// 	defer span.End()

// 	req := &productsv1.CreateProductRequest{
// 		Name:        "my amazing",
// 		ImageUrl:    "https://amazing.life/image.png",
// 		Description: "this really is an amazing product you would not imagine",
// 		Price:       9.99,
// 		Category:    productsv1.ProductCategory_PRODUCT_CATEGORY_BOOK,
// 	}

// 	resp, err := productsClient.CreateProduct(ctx, connect.NewRequest(req))
// 	if err != nil {
// 		return fmt.Errorf("failed to create product: %w", err)
// 	}
// 	lgr.Info("got resp from creating product", "resp", resp.Msg.GetProduct().GetId())

// 	createReviewReq := &reviewsv1.CreateReviewRequest{
// 		ProductId: resp.Msg.GetProduct().GetId(),
// 		Title:     "this is my amazing review",
// 		Body:      "this product really is a life saver i do not know what i would do without it",
// 		Rating:    5,
// 	}

// 	createReviewResp, err := reviewsClient.CreateReview(ctx, connect.NewRequest(createReviewReq))
// 	if err != nil {
// 		return fmt.Errorf("failed to create review for %s: %w", createReviewReq.GetProductId(), err)
// 	}

// 	lgr.Info("got resp from creating review", "resp", createReviewResp.Msg.GetReview())

// 	// overviewReq := &productsv1.GetProductOverviewRequest{ProductId: createReviewReq.GetProductId(), FieldMask: &fieldmaskpb.FieldMask{Paths: []string{"reviews"}}}
// 	// overviewResp, err := productsClient.GetProductOverview(ctx, overviewReq)
// 	// if err != nil {
// 	// 	return fmt.Errorf("failed to create review for %s: %w", createReviewReq.GetProductId(), err)
// 	// }

// 	// lgr.Info("got resp from overview review", "resp", overviewResp.GetProductOverview())

// 	return nil
// }
