package main

//func TestE2E(t *testing.T) {
//	go startSrv()
//	client := getClient(t)
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//
//}
//
//func getClient(t *testing.T) productsv1.ProductsServiceClient {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//	conn, err := grpc.DialContext(ctx, "localhost:50051",
//		grpc.WithInsecure(),
//		grpc.WithPerRPCCredentials(xgrpc.TokenAuth{
//			Token: "ciccio",
//		}),
//	)
//	require.NoError(t, err)
//
//	return productsv1.NewProductsServiceClient(conn)
//}
//
//func startSrv() {
//	lgr := logrus.New()
//	srv, err := web.New(logrus.NewEntry(lgr), ":50051")
//	if err != nil {
//		logrus.WithError(err).Fatal("failed to listen on port")
//	}
//	if err := srv.Run(); err != nil {
//		logrus.WithError(err).Fatal("failed to run server")
//	}
//}
