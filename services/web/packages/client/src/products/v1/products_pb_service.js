// package: products.v1
// file: products/v1/products.proto

var products_v1_products_pb = require("../../products/v1/products_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var ProductsService = (function () {
  function ProductsService() {}
  ProductsService.serviceName = "products.v1.ProductsService";
  return ProductsService;
}());

ProductsService.ListProducts = {
  methodName: "ListProducts",
  service: ProductsService,
  requestStream: false,
  responseStream: false,
  requestType: products_v1_products_pb.ListProductsRequest,
  responseType: products_v1_products_pb.ListProductsResponse
};

ProductsService.CreateProduct = {
  methodName: "CreateProduct",
  service: ProductsService,
  requestStream: false,
  responseStream: false,
  requestType: products_v1_products_pb.CreateProductRequest,
  responseType: products_v1_products_pb.CreateProductResponse
};

ProductsService.UpdateProduct = {
  methodName: "UpdateProduct",
  service: ProductsService,
  requestStream: false,
  responseStream: false,
  requestType: products_v1_products_pb.UpdateProductRequest,
  responseType: products_v1_products_pb.UpdateProductResponse
};

ProductsService.GetProduct = {
  methodName: "GetProduct",
  service: ProductsService,
  requestStream: false,
  responseStream: false,
  requestType: products_v1_products_pb.GetProductRequest,
  responseType: products_v1_products_pb.GetProductResponse
};

ProductsService.GetProductOverview = {
  methodName: "GetProductOverview",
  service: ProductsService,
  requestStream: false,
  responseStream: false,
  requestType: products_v1_products_pb.GetProductOverviewRequest,
  responseType: products_v1_products_pb.GetProductOverviewResponse
};

exports.ProductsService = ProductsService;

function ProductsServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ProductsServiceClient.prototype.listProducts = function listProducts(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ProductsService.ListProducts, {
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

ProductsServiceClient.prototype.createProduct = function createProduct(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ProductsService.CreateProduct, {
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

ProductsServiceClient.prototype.updateProduct = function updateProduct(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ProductsService.UpdateProduct, {
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

ProductsServiceClient.prototype.getProduct = function getProduct(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ProductsService.GetProduct, {
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

ProductsServiceClient.prototype.getProductOverview = function getProductOverview(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ProductsService.GetProductOverview, {
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

exports.ProductsServiceClient = ProductsServiceClient;

