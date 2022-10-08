// @generated by protoc-gen-connect-web v0.2.1 with parameter "target=js+dts"
// @generated from file products/v1/products.proto (package products.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import {CreateProductRequest, CreateProductResponse, GetProductOverviewRequest, GetProductOverviewResponse, GetProductRequest, GetProductResponse, ListProductsRequest, ListProductsResponse, UpdateProductRequest, UpdateProductResponse} from "./products_pb.js";
import {MethodKind} from "@bufbuild/protobuf";

/**
 * @generated from service products.v1.ProductsService
 */
export declare const ProductsService: {
  readonly typeName: "products.v1.ProductsService",
  readonly methods: {
    /**
     * @generated from rpc products.v1.ProductsService.ListProducts
     */
    readonly listProducts: {
      readonly name: "ListProducts",
      readonly I: typeof ListProductsRequest,
      readonly O: typeof ListProductsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc products.v1.ProductsService.CreateProduct
     */
    readonly createProduct: {
      readonly name: "CreateProduct",
      readonly I: typeof CreateProductRequest,
      readonly O: typeof CreateProductResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc products.v1.ProductsService.UpdateProduct
     */
    readonly updateProduct: {
      readonly name: "UpdateProduct",
      readonly I: typeof UpdateProductRequest,
      readonly O: typeof UpdateProductResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc products.v1.ProductsService.GetProduct
     */
    readonly getProduct: {
      readonly name: "GetProduct",
      readonly I: typeof GetProductRequest,
      readonly O: typeof GetProductResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc products.v1.ProductsService.GetProductOverview
     */
    readonly getProductOverview: {
      readonly name: "GetProductOverview",
      readonly I: typeof GetProductOverviewRequest,
      readonly O: typeof GetProductOverviewResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};
