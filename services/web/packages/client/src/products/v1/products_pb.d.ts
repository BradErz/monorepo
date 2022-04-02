// package: products.v1
// file: products/v1/products.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_field_mask_pb from "google-protobuf/google/protobuf/field_mask_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as reviews_v1_reviews_pb from "../../reviews/v1/reviews_pb";

export class GetProductOverviewRequest extends jspb.Message {
  getProductId(): string;
  setProductId(value: string): void;

  hasFieldMask(): boolean;
  clearFieldMask(): void;
  getFieldMask(): google_protobuf_field_mask_pb.FieldMask | undefined;
  setFieldMask(value?: google_protobuf_field_mask_pb.FieldMask): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetProductOverviewRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetProductOverviewRequest): GetProductOverviewRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetProductOverviewRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetProductOverviewRequest;
  static deserializeBinaryFromReader(message: GetProductOverviewRequest, reader: jspb.BinaryReader): GetProductOverviewRequest;
}

export namespace GetProductOverviewRequest {
  export type AsObject = {
    productId: string,
    fieldMask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
  }
}

export class GetProductOverviewResponse extends jspb.Message {
  hasProductOverview(): boolean;
  clearProductOverview(): void;
  getProductOverview(): ProductOverview | undefined;
  setProductOverview(value?: ProductOverview): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetProductOverviewResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetProductOverviewResponse): GetProductOverviewResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetProductOverviewResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetProductOverviewResponse;
  static deserializeBinaryFromReader(message: GetProductOverviewResponse, reader: jspb.BinaryReader): GetProductOverviewResponse;
}

export namespace GetProductOverviewResponse {
  export type AsObject = {
    productOverview?: ProductOverview.AsObject,
  }
}

export class ProductOverview extends jspb.Message {
  hasProduct(): boolean;
  clearProduct(): void;
  getProduct(): Product | undefined;
  setProduct(value?: Product): void;

  clearReviewsList(): void;
  getReviewsList(): Array<reviews_v1_reviews_pb.Review>;
  setReviewsList(value: Array<reviews_v1_reviews_pb.Review>): void;
  addReviews(value?: reviews_v1_reviews_pb.Review, index?: number): reviews_v1_reviews_pb.Review;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ProductOverview.AsObject;
  static toObject(includeInstance: boolean, msg: ProductOverview): ProductOverview.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ProductOverview, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ProductOverview;
  static deserializeBinaryFromReader(message: ProductOverview, reader: jspb.BinaryReader): ProductOverview;
}

export namespace ProductOverview {
  export type AsObject = {
    product?: Product.AsObject,
    reviewsList: Array<reviews_v1_reviews_pb.Review.AsObject>,
  }
}

export class GetProductRequest extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetProductRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetProductRequest): GetProductRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetProductRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetProductRequest;
  static deserializeBinaryFromReader(message: GetProductRequest, reader: jspb.BinaryReader): GetProductRequest;
}

export namespace GetProductRequest {
  export type AsObject = {
    id: string,
  }
}

export class GetProductResponse extends jspb.Message {
  hasProduct(): boolean;
  clearProduct(): void;
  getProduct(): Product | undefined;
  setProduct(value?: Product): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetProductResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetProductResponse): GetProductResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetProductResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetProductResponse;
  static deserializeBinaryFromReader(message: GetProductResponse, reader: jspb.BinaryReader): GetProductResponse;
}

export namespace GetProductResponse {
  export type AsObject = {
    product?: Product.AsObject,
  }
}

export class ListProductsRequest extends jspb.Message {
  getPageSize(): number;
  setPageSize(value: number): void;

  getPageToken(): string;
  setPageToken(value: string): void;

  getFilter(): string;
  setFilter(value: string): void;

  getOrderBy(): string;
  setOrderBy(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListProductsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListProductsRequest): ListProductsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListProductsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListProductsRequest;
  static deserializeBinaryFromReader(message: ListProductsRequest, reader: jspb.BinaryReader): ListProductsRequest;
}

export namespace ListProductsRequest {
  export type AsObject = {
    pageSize: number,
    pageToken: string,
    filter: string,
    orderBy: string,
  }
}

export class ListProductsResponse extends jspb.Message {
  clearProductsList(): void;
  getProductsList(): Array<Product>;
  setProductsList(value: Array<Product>): void;
  addProducts(value?: Product, index?: number): Product;

  getNextPageToken(): string;
  setNextPageToken(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListProductsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListProductsResponse): ListProductsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListProductsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListProductsResponse;
  static deserializeBinaryFromReader(message: ListProductsResponse, reader: jspb.BinaryReader): ListProductsResponse;
}

export namespace ListProductsResponse {
  export type AsObject = {
    productsList: Array<Product.AsObject>,
    nextPageToken: string,
  }
}

export class CreateProductRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getImageUrl(): string;
  setImageUrl(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  getPrice(): number;
  setPrice(value: number): void;

  getCategory(): ProductCategoryMap[keyof ProductCategoryMap];
  setCategory(value: ProductCategoryMap[keyof ProductCategoryMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateProductRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateProductRequest): CreateProductRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateProductRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateProductRequest;
  static deserializeBinaryFromReader(message: CreateProductRequest, reader: jspb.BinaryReader): CreateProductRequest;
}

export namespace CreateProductRequest {
  export type AsObject = {
    name: string,
    imageUrl: string,
    description: string,
    price: number,
    category: ProductCategoryMap[keyof ProductCategoryMap],
  }
}

export class CreateProductResponse extends jspb.Message {
  hasProduct(): boolean;
  clearProduct(): void;
  getProduct(): Product | undefined;
  setProduct(value?: Product): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateProductResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateProductResponse): CreateProductResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateProductResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateProductResponse;
  static deserializeBinaryFromReader(message: CreateProductResponse, reader: jspb.BinaryReader): CreateProductResponse;
}

export namespace CreateProductResponse {
  export type AsObject = {
    product?: Product.AsObject,
  }
}

export class UpdateProductRequest extends jspb.Message {
  hasProduct(): boolean;
  clearProduct(): void;
  getProduct(): Product | undefined;
  setProduct(value?: Product): void;

  hasFieldMask(): boolean;
  clearFieldMask(): void;
  getFieldMask(): google_protobuf_field_mask_pb.FieldMask | undefined;
  setFieldMask(value?: google_protobuf_field_mask_pb.FieldMask): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateProductRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateProductRequest): UpdateProductRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateProductRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateProductRequest;
  static deserializeBinaryFromReader(message: UpdateProductRequest, reader: jspb.BinaryReader): UpdateProductRequest;
}

export namespace UpdateProductRequest {
  export type AsObject = {
    product?: Product.AsObject,
    fieldMask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
  }
}

export class UpdateProductResponse extends jspb.Message {
  hasProduct(): boolean;
  clearProduct(): void;
  getProduct(): Product | undefined;
  setProduct(value?: Product): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateProductResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateProductResponse): UpdateProductResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpdateProductResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateProductResponse;
  static deserializeBinaryFromReader(message: UpdateProductResponse, reader: jspb.BinaryReader): UpdateProductResponse;
}

export namespace UpdateProductResponse {
  export type AsObject = {
    product?: Product.AsObject,
  }
}

export class Product extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  hasCreateTime(): boolean;
  clearCreateTime(): void;
  getCreateTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreateTime(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasUpdateTime(): boolean;
  clearUpdateTime(): void;
  getUpdateTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdateTime(value?: google_protobuf_timestamp_pb.Timestamp): void;

  hasDeleteTime(): boolean;
  clearDeleteTime(): void;
  getDeleteTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDeleteTime(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getName(): string;
  setName(value: string): void;

  getImageUrl(): string;
  setImageUrl(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  getPrice(): number;
  setPrice(value: number): void;

  getCategory(): ProductCategoryMap[keyof ProductCategoryMap];
  setCategory(value: ProductCategoryMap[keyof ProductCategoryMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Product.AsObject;
  static toObject(includeInstance: boolean, msg: Product): Product.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Product, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Product;
  static deserializeBinaryFromReader(message: Product, reader: jspb.BinaryReader): Product;
}

export namespace Product {
  export type AsObject = {
    id: string,
    createTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updateTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    deleteTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    name: string,
    imageUrl: string,
    description: string,
    price: number,
    category: ProductCategoryMap[keyof ProductCategoryMap],
  }
}

export interface ProductCategoryMap {
  PRODUCT_CATEGORY_UNSPECIFIED: 0;
  PRODUCT_CATEGORY_SOFTWARE: 1;
  PRODUCT_CATEGORY_BOOK: 2;
  PRODUCT_CATEGORY_COURSE: 3;
}

export const ProductCategory: ProductCategoryMap;

