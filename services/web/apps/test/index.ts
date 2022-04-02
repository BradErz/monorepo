import {ProductsServiceClient, GetProductOverviewRequest} from "client";
import * as google_protobuf_field_mask_pb from "google-protobuf/google/protobuf/field_mask_pb";


const productsClient = new ProductsServiceClient("http://localhost:10000")
const overviewRequest = new GetProductOverviewRequest()
overviewRequest.setProductId("624755cd92e2ce89184fe0cf")
const value10 = new google_protobuf_field_mask_pb.FieldMask;
value10.setPathsList(["reviews"])
overviewRequest.setFieldMask(value10)

productsClient.getProductOverview(overviewRequest, (err, resp) => {
    if (err !== null) {
        console.log(`error code ${err.code} message ${err.message}`)
        return
    }

    // @ts-ignore
    const product = resp.getProductOverview().getProduct()
    // @ts-ignore
    console.log(`product ${product.getName()} category: ${product.getCategory()} price: ${product.getPrice()}`)
    // @ts-ignore

    const reviews = resp.getProductOverview().getReviewsList()
    for (const reviewsListKey in reviews) {
        console.log(`got review ${reviews[reviewsListKey].getTitle()} with rating ${reviews[reviewsListKey].getRating()}`)
    }
});
