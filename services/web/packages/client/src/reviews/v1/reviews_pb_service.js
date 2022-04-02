// package: reviews.v1
// file: reviews/v1/reviews.proto

var reviews_v1_reviews_pb = require("../../reviews/v1/reviews_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var ReviewsService = (function () {
  function ReviewsService() {}
  ReviewsService.serviceName = "reviews.v1.ReviewsService";
  return ReviewsService;
}());

ReviewsService.CreateReview = {
  methodName: "CreateReview",
  service: ReviewsService,
  requestStream: false,
  responseStream: false,
  requestType: reviews_v1_reviews_pb.CreateReviewRequest,
  responseType: reviews_v1_reviews_pb.CreateReviewResponse
};

ReviewsService.ListReviews = {
  methodName: "ListReviews",
  service: ReviewsService,
  requestStream: false,
  responseStream: false,
  requestType: reviews_v1_reviews_pb.ListReviewsRequest,
  responseType: reviews_v1_reviews_pb.ListReviewsResponse
};

exports.ReviewsService = ReviewsService;

function ReviewsServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ReviewsServiceClient.prototype.createReview = function createReview(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ReviewsService.CreateReview, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ReviewsServiceClient.prototype.listReviews = function listReviews(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ReviewsService.ListReviews, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.ReviewsServiceClient = ReviewsServiceClient;

