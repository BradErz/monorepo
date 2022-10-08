package web

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/go-logr/logr"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/BradErz/monorepo/services/reviews/models"

	reviewsv1 "github.com/BradErz/monorepo/gen/go/reviews/v1"
	"github.com/BradErz/monorepo/gen/go/reviews/v1/reviewsv1connect"
)

type Server struct {
	lgr     logr.Logger
	service Service
}

var _ reviewsv1connect.ReviewsServiceHandler = (*Server)(nil)

func New(lgr logr.Logger, service Service) *Server {
	return &Server{
		lgr:     lgr.WithName("reviewsv1"),
		service: service,
	}
}

func (s *Server) CreateReview(ctx context.Context, req *connect.Request[reviewsv1.CreateReviewRequest]) (*connect.Response[reviewsv1.CreateReviewResponse], error) {
	resp, err := s.service.CreateReview(ctx, toModelCreateReviewReq(req.Msg))
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&reviewsv1.CreateReviewResponse{
		Review: toProtoReview(resp),
	}), nil
}

func (s *Server) ListReviews(ctx context.Context, req *connect.Request[reviewsv1.ListReviewsRequest]) (*connect.Response[reviewsv1.ListReviewsResponse], error) {
	resp, err := s.service.ListReviews(ctx, toModelListReviewReq(req.Msg))
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(toProtoListReviewResponse(resp)), nil
}

func Register(svc reviewsv1connect.ReviewsServiceHandler, opts ...connect.HandlerOption) {
}

// func Register(reviewSrv reviewsv1connect.ReviewsServiceHandler) xgrpc.RegisterServerFunc {
// 	return func(s *grpc.Server) {
// 		reviewsv1connect.RegisterReviewsServiceServer(s, reviewSrv)
// 	}
// }

// func (srv *Server) CreateReview(ctx context.Context, req *reviewsv1.CreateReviewRequest) (*reviewsv1.CreateReviewResponse, error) {
// 	// if err := req.ValidateAll(); err != nil {
// 	// 	srv.lgr.Error(err, "request was not valid")
// 	// 	return nil, xerrors.Wrapf(xerrors.CodeInvalidArgument, err, "request was invalid: %s", err.Error())
// 	// }

// 	resp, err := srv.service.CreateReview(ctx, toModelCreateReviewReq(req))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &reviewsv1.CreateReviewResponse{Review: toProtoReview(resp)}, nil
// }

// func (srv *Server) ListReviews(ctx context.Context, req *reviewsv1.ListReviewsRequest) (*reviewsv1.ListReviewsResponse, error) {
// 	resp, err := srv.service.ListReviews(ctx, toModelListReviewReq(req))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return toProtoListReviewResponse(resp), nil
// }

func toModelListReviewReq(req *reviewsv1.ListReviewsRequest) *models.ListReviewsRequest {
	return &models.ListReviewsRequest{
		ProductID: req.GetProductId(),
		PageSize:  req.GetPageSize(),
		PageToken: req.GetPageToken(),
	}
}

func toModelCreateReviewReq(req *reviewsv1.CreateReviewRequest) *models.CreateReviewRequest {
	return &models.CreateReviewRequest{
		ProductID: req.GetProductId(),
		Title:     req.GetTitle(),
		Body:      req.GetBody(),
		Rating:    uint(req.GetRating()),
	}
}

func toProtoListReviewResponse(resp *models.ListReviewsResponse) *reviewsv1.ListReviewsResponse {
	reviews := make([]*reviewsv1.Review, len(resp.Reviews))
	for i, review := range resp.Reviews {
		reviews[i] = toProtoReview(review)
	}

	return &reviewsv1.ListReviewsResponse{
		Reviews:       reviews,
		NextPageToken: resp.NextPageToken,
	}
}

func toProtoReview(review *models.Review) *reviewsv1.Review {
	r := &reviewsv1.Review{
		Id:         review.ID,
		ProductId:  review.ProductID,
		CreateTime: timestamppb.New(review.CreateTime),
		Title:      review.Title,
		Body:       review.Body,
		Rating:     uint32(review.Rating),
	}
	if review.UpdateTime != nil {
		r.UpdateTime = timestamppb.New(*review.UpdateTime)
	}
	if review.DeleteTime != nil {
		r.UpdateTime = timestamppb.New(*review.DeleteTime)
	}
	return r
}
