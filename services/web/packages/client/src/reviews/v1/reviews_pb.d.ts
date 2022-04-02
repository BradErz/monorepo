// package: reviews.v1
// file: reviews/v1/reviews.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class CreateReviewRequest extends jspb.Message {
  getProductId(): string;
  setProductId(value: string): void;

  getTitle(): string;
  setTitle(value: string): void;

  getBody(): string;
  setBody(value: string): void;

  getRating(): number;
  setRating(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateReviewRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateReviewRequest): CreateReviewRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateReviewRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateReviewRequest;
  static deserializeBinaryFromReader(message: CreateReviewRequest, reader: jspb.BinaryReader): CreateReviewRequest;
}

export namespace CreateReviewRequest {
  export type AsObject = {
    productId: string,
    title: string,
    body: string,
    rating: number,
  }
}

export class CreateReviewResponse extends jspb.Message {
  hasReview(): boolean;
  clearReview(): void;
  getReview(): Review | undefined;
  setReview(value?: Review): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateReviewResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateReviewResponse): CreateReviewResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateReviewResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateReviewResponse;
  static deserializeBinaryFromReader(message: CreateReviewResponse, reader: jspb.BinaryReader): CreateReviewResponse;
}

export namespace CreateReviewResponse {
  export type AsObject = {
    review?: Review.AsObject,
  }
}

export class ListReviewsRequest extends jspb.Message {
  getProductId(): string;
  setProductId(value: string): void;

  getPageSize(): number;
  setPageSize(value: number): void;

  getPageToken(): string;
  setPageToken(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListReviewsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListReviewsRequest): ListReviewsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListReviewsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListReviewsRequest;
  static deserializeBinaryFromReader(message: ListReviewsRequest, reader: jspb.BinaryReader): ListReviewsRequest;
}

export namespace ListReviewsRequest {
  export type AsObject = {
    productId: string,
    pageSize: number,
    pageToken: string,
  }
}

export class ListReviewsResponse extends jspb.Message {
  clearReviewsList(): void;
  getReviewsList(): Array<Review>;
  setReviewsList(value: Array<Review>): void;
  addReviews(value?: Review, index?: number): Review;

  getNextPageToken(): string;
  setNextPageToken(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListReviewsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListReviewsResponse): ListReviewsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListReviewsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListReviewsResponse;
  static deserializeBinaryFromReader(message: ListReviewsResponse, reader: jspb.BinaryReader): ListReviewsResponse;
}

export namespace ListReviewsResponse {
  export type AsObject = {
    reviewsList: Array<Review.AsObject>,
    nextPageToken: string,
  }
}

export class Review extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getProductId(): string;
  setProductId(value: string): void;

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

  getTitle(): string;
  setTitle(value: string): void;

  getBody(): string;
  setBody(value: string): void;

  getRating(): number;
  setRating(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Review.AsObject;
  static toObject(includeInstance: boolean, msg: Review): Review.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Review, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Review;
  static deserializeBinaryFromReader(message: Review, reader: jspb.BinaryReader): Review;
}

export namespace Review {
  export type AsObject = {
    id: string,
    productId: string,
    createTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updateTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    deleteTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    title: string,
    body: string,
    rating: number,
  }
}

