// @generated by protoc-gen-es v0.1.1 with parameter "target=js+dts"
// @generated from file products/v1/products.proto (package products.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import {FieldMask, proto3, Timestamp} from "@bufbuild/protobuf";
import {Review} from "../../reviews/v1/reviews_pb.js";

/**
 * @generated from enum products.v1.ProductCategory
 */
export const ProductCategory = proto3.makeEnum(
  "products.v1.ProductCategory",
  [
    {no: 0, name: "PRODUCT_CATEGORY_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "PRODUCT_CATEGORY_SOFTWARE", localName: "SOFTWARE"},
    {no: 2, name: "PRODUCT_CATEGORY_BOOK", localName: "BOOK"},
    {no: 3, name: "PRODUCT_CATEGORY_COURSE", localName: "COURSE"},
  ],
);

/**
 * @generated from message products.v1.GetProductOverviewRequest
 */
export const GetProductOverviewRequest = proto3.makeMessageType(
  "products.v1.GetProductOverviewRequest",
  () => [
    { no: 1, name: "product_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "field_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message products.v1.GetProductOverviewResponse
 */
export const GetProductOverviewResponse = proto3.makeMessageType(
  "products.v1.GetProductOverviewResponse",
  () => [
    { no: 1, name: "product_overview", kind: "message", T: ProductOverview },
  ],
);

/**
 * @generated from message products.v1.ProductOverview
 */
export const ProductOverview = proto3.makeMessageType(
  "products.v1.ProductOverview",
  () => [
    { no: 1, name: "product", kind: "message", T: Product },
    { no: 2, name: "reviews", kind: "message", T: Review, repeated: true },
  ],
);

/**
 * @generated from message products.v1.GetProductRequest
 */
export const GetProductRequest = proto3.makeMessageType(
  "products.v1.GetProductRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message products.v1.GetProductResponse
 */
export const GetProductResponse = proto3.makeMessageType(
  "products.v1.GetProductResponse",
  () => [
    { no: 1, name: "product", kind: "message", T: Product },
  ],
);

/**
 * @generated from message products.v1.ListProductsRequest
 */
export const ListProductsRequest = proto3.makeMessageType(
  "products.v1.ListProductsRequest",
  () => [
    { no: 1, name: "page_size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "filter", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "order_by", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message products.v1.ListProductsResponse
 */
export const ListProductsResponse = proto3.makeMessageType(
  "products.v1.ListProductsResponse",
  () => [
    { no: 1, name: "products", kind: "message", T: Product, repeated: true },
    { no: 2, name: "next_page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message products.v1.CreateProductRequest
 */
export const CreateProductRequest = proto3.makeMessageType(
  "products.v1.CreateProductRequest",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "image_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "price", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 5, name: "category", kind: "enum", T: proto3.getEnumType(ProductCategory) },
  ],
);

/**
 * @generated from message products.v1.CreateProductResponse
 */
export const CreateProductResponse = proto3.makeMessageType(
  "products.v1.CreateProductResponse",
  () => [
    { no: 1, name: "product", kind: "message", T: Product },
  ],
);

/**
 * @generated from message products.v1.UpdateProductRequest
 */
export const UpdateProductRequest = proto3.makeMessageType(
  "products.v1.UpdateProductRequest",
  () => [
    { no: 1, name: "product", kind: "message", T: Product },
    { no: 2, name: "field_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message products.v1.UpdateProductResponse
 */
export const UpdateProductResponse = proto3.makeMessageType(
  "products.v1.UpdateProductResponse",
  () => [
    { no: 1, name: "product", kind: "message", T: Product },
  ],
);

/**
 * @generated from message products.v1.Product
 */
export const Product = proto3.makeMessageType(
  "products.v1.Product",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "create_time", kind: "message", T: Timestamp },
    { no: 3, name: "update_time", kind: "message", T: Timestamp },
    { no: 4, name: "delete_time", kind: "message", T: Timestamp },
    { no: 10, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 11, name: "image_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 12, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 13, name: "price", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 14, name: "category", kind: "enum", T: proto3.getEnumType(ProductCategory) },
  ],
);
