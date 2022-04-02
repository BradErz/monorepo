import {Button} from "ui";
import {ProductsServiceClient, GetProductOverviewRequest} from "client";
import {grpc} from "@improbable-eng/grpc-web";


const apiEndpoint = process.env.NEXT_PUBLIC_API_ENDPOINT

export default function Docs() {
    const productsClient = new ProductsServiceClient(apiEndpoint);
    const overviewReq = new GetProductOverviewRequest();
    overviewReq.setProductId("624755cd92e2ce89184fe0cf");

    productsClient.getProductOverview(overviewReq, (err, resp) => {
        console.log(`error code ${err.code} message ${err.message}`)
        console.log(`response: ${resp.getProductOverview()}`)
    });

    return (
        <div>
            <h1>Docs</h1>
            <Button/>
        </div>
    );
}
