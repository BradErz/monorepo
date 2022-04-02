// package: reviews.v1
// file: reviews/v1/reviews.proto

import * as reviews_v1_reviews_pb from "../../reviews/v1/reviews_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ReviewsServiceCreateReview = {
  readonly methodName: string;
  readonly service: typeof ReviewsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof reviews_v1_reviews_pb.CreateReviewRequest;
  readonly responseType: typeof reviews_v1_reviews_pb.CreateReviewResponse;
};

type ReviewsServiceListReviews = {
  readonly methodName: string;
  readonly service: typeof ReviewsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof reviews_v1_reviews_pb.ListReviewsRequest;
  readonly responseType: typeof reviews_v1_reviews_pb.ListReviewsResponse;
};

export class ReviewsService {
  static readonly serviceName: string;
  static readonly CreateReview: ReviewsServiceCreateReview;
  static readonly ListReviews: ReviewsServiceListReviews;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class ReviewsServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  createReview(
    requestMessage: reviews_v1_reviews_pb.CreateReviewRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: reviews_v1_reviews_pb.CreateReviewResponse|null) => void
  ): UnaryResponse;
  createReview(
    requestMessage: reviews_v1_reviews_pb.CreateReviewRequest,
    callback: (error: ServiceError|null, responseMessage: reviews_v1_reviews_pb.CreateReviewResponse|null) => void
  ): UnaryResponse;
  listReviews(
    requestMessage: reviews_v1_reviews_pb.ListReviewsRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: reviews_v1_reviews_pb.ListReviewsResponse|null) => void
  ): UnaryResponse;
  listReviews(
    requestMessage: reviews_v1_reviews_pb.ListReviewsRequest,
    callback: (error: ServiceError|null, responseMessage: reviews_v1_reviews_pb.ListReviewsResponse|null) => void
  ): UnaryResponse;
}

