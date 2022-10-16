package service

import (
	"context"
	"fmt"

	"github.com/BradErz/monorepo/pkg/xerrors"
	"github.com/BradErz/monorepo/services/reviews/models"
	"github.com/bufbuild/connect-go"
	productsv1 "go.buf.build/bufbuild/connect-go/braderz/monorepo/products/v1"
	productsv1connect "go.buf.build/bufbuild/connect-go/braderz/monorepo/products/v1/productsv1connect"
)

type Reviews struct {
	store            Storage
	productsV1Client productsv1connect.ProductsServiceClient
}

func NewReviews(store Storage, productsV1Client productsv1connect.ProductsServiceClient) *Reviews {
	return &Reviews{
		store:            store,
		productsV1Client: productsV1Client,
	}
}

func (r *Reviews) CreateReview(ctx context.Context, req *models.CreateReviewRequest) (*models.Review, error) {
	_, err := r.productsV1Client.GetProduct(ctx, connect.NewRequest(&productsv1.GetProductRequest{Id: req.ProductID}))
	switch {
	case connect.CodeOf(err) == connect.CodeNotFound:
		return nil, xerrors.Wrapf(xerrors.CodeNotFound, err, "product %s doesnt exist in the products service", req.ProductID)
	case connect.CodeOf(err) == connect.CodeInvalidArgument:
		return nil, xerrors.Wrapf(xerrors.CodeInvalidArgument, err, "product %s was invalid: %s", req.ProductID, err)
	case err != nil:
		fmt.Println(err)
		return nil, xerrors.Wrapf(xerrors.CodeInternal, err, "could not call products service to check if %s existed", req.ProductID)
	}

	product, err := r.store.CreateReview(ctx, req)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *Reviews) ListReviews(ctx context.Context, req *models.ListReviewsRequest) (*models.ListReviewsResponse, error) {
	listResp, err := r.store.ListReviews(ctx, req)
	if err != nil {
		return nil, err
	}
	return listResp, nil
}
