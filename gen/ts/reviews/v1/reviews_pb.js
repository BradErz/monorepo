// @generated by protoc-gen-es v0.1.1 with parameter "target=js+dts"
// @generated from file reviews/v1/reviews.proto (package reviews.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import {proto3, Timestamp} from "@bufbuild/protobuf";

/**
 * @generated from message reviews.v1.CreateReviewRequest
 */
export const CreateReviewRequest = proto3.makeMessageType(
  "reviews.v1.CreateReviewRequest",
  () => [
    { no: 1, name: "product_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "body", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "rating", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
  ],
);

/**
 * @generated from message reviews.v1.CreateReviewResponse
 */
export const CreateReviewResponse = proto3.makeMessageType(
  "reviews.v1.CreateReviewResponse",
  () => [
    { no: 1, name: "review", kind: "message", T: Review },
  ],
);

/**
 * @generated from message reviews.v1.ListReviewsRequest
 */
export const ListReviewsRequest = proto3.makeMessageType(
  "reviews.v1.ListReviewsRequest",
  () => [
    { no: 1, name: "product_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "page_size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message reviews.v1.ListReviewsResponse
 */
export const ListReviewsResponse = proto3.makeMessageType(
  "reviews.v1.ListReviewsResponse",
  () => [
    { no: 1, name: "reviews", kind: "message", T: Review, repeated: true },
    { no: 2, name: "next_page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message reviews.v1.Review
 */
export const Review = proto3.makeMessageType(
  "reviews.v1.Review",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "product_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "create_time", kind: "message", T: Timestamp },
    { no: 4, name: "update_time", kind: "message", T: Timestamp },
    { no: 5, name: "delete_time", kind: "message", T: Timestamp },
    { no: 10, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 11, name: "body", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 12, name: "rating", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
  ],
);

