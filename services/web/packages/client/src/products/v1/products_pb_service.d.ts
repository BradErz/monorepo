// package: products.v1
// file: products/v1/products.proto

import * as products_v1_products_pb from "../../products/v1/products_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ProductsServiceListProducts = {
  readonly methodName: string;
  readonly service: typeof ProductsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof products_v1_products_pb.ListProductsRequest;
  readonly responseType: typeof products_v1_products_pb.ListProductsResponse;
};

type ProductsServiceCreateProduct = {
  readonly methodName: string;
  readonly service: typeof ProductsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof products_v1_products_pb.CreateProductRequest;
  readonly responseType: typeof products_v1_products_pb.CreateProductResponse;
};

type ProductsServiceUpdateProduct = {
  readonly methodName: string;
  readonly service: typeof ProductsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof products_v1_products_pb.UpdateProductRequest;
  readonly responseType: typeof products_v1_products_pb.UpdateProductResponse;
};

type ProductsServiceGetProduct = {
  readonly methodName: string;
  readonly service: typeof ProductsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof products_v1_products_pb.GetProductRequest;
  readonly responseType: typeof products_v1_products_pb.GetProductResponse;
};

type ProductsServiceGetProductOverview = {
  readonly methodName: string;
  readonly service: typeof ProductsService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof products_v1_products_pb.GetProductOverviewRequest;
  readonly responseType: typeof products_v1_products_pb.GetProductOverviewResponse;
};

export class ProductsService {
  static readonly serviceName: string;
  static readonly ListProducts: ProductsServiceListProducts;
  static readonly CreateProduct: ProductsServiceCreateProduct;
  static readonly UpdateProduct: ProductsServiceUpdateProduct;
  static readonly GetProduct: ProductsServiceGetProduct;
  static readonly GetProductOverview: ProductsServiceGetProductOverview;
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

export class ProductsServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  listProducts(
    requestMessage: products_v1_products_pb.ListProductsRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: products_v1_products_pb.ListProductsResponse|null) => void
  ): UnaryResponse;
  listProducts(
    requestMessage: products_v1_products_pb.ListProductsRequest,
    callback: (error: ServiceError|null, responseMessage: products_v1_products_pb.ListProductsResponse|null) => void
  ): UnaryResponse;
  createProduct(
    requestMessage: products_v1_products_pb.CreateProductRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: products_v1_products_pb.CreateProductResponse|null) => void
  ): UnaryResponse;
  createProduct(
    requestMessage: products_v1_products_pb.CreateProductRequest,
    callback: (error: ServiceError|null, responseMessage: products_v1_products_pb.CreateProductResponse|null) => void
  ): UnaryResponse;
  updateProduct(
    requestMessage: products_v1_products_pb.UpdateProductRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: products_v1_products_pb.UpdateProductResponse|null) => void
  ): UnaryResponse;
  updateProduct(
    requestMessage: products_v1_products_pb.UpdateProductRequest,
    callback: (error: ServiceError|null, responseMessage: products_v1_products_pb.UpdateProductResponse|null) => void
  ): UnaryResponse;
  getProduct(
    requestMessage: products_v1_products_pb.GetProductRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: products_v1_products_pb.GetProductResponse|null) => void
  ): UnaryResponse;
  getProduct(
    requestMessage: products_v1_products_pb.GetProductRequest,
    callback: (error: ServiceError|null, responseMessage: products_v1_products_pb.GetProductResponse|null) => void
  ): UnaryResponse;
  getProductOverview(
    requestMessage: products_v1_products_pb.GetProductOverviewRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: products_v1_products_pb.GetProductOverviewResponse|null) => void
  ): UnaryResponse;
  getProductOverview(
    requestMessage: products_v1_products_pb.GetProductOverviewRequest,
    callback: (error: ServiceError|null, responseMessage: products_v1_products_pb.GetProductOverviewResponse|null) => void
  ): UnaryResponse;
}

