// @generated by protoc-gen-connect-web v0.2.1 with parameter "target=ts"
// @generated from file products/v1/products.proto (package products.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import {CreateProductRequest, CreateProductResponse, GetProductOverviewRequest, GetProductOverviewResponse, GetProductRequest, GetProductResponse, ListProductsRequest, ListProductsResponse, UpdateProductRequest, UpdateProductResponse} from "./products_pb.js";
import {MethodKind} from "@bufbuild/protobuf";

/**
 * @generated from service products.v1.ProductsService
 */
export const ProductsService = {
  typeName: "products.v1.ProductsService",
  methods: {
    /**
     * @generated from rpc products.v1.ProductsService.ListProducts
     */
    listProducts: {
      name: "ListProducts",
      I: ListProductsRequest,
      O: ListProductsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc products.v1.ProductsService.CreateProduct
     */
    createProduct: {
      name: "CreateProduct",
      I: CreateProductRequest,
      O: CreateProductResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc products.v1.ProductsService.UpdateProduct
     */
    updateProduct: {
      name: "UpdateProduct",
      I: UpdateProductRequest,
      O: UpdateProductResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc products.v1.ProductsService.GetProduct
     */
    getProduct: {
      name: "GetProduct",
      I: GetProductRequest,
      O: GetProductResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc products.v1.ProductsService.GetProductOverview
     */
    getProductOverview: {
      name: "GetProductOverview",
      I: GetProductOverviewRequest,
      O: GetProductOverviewResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

