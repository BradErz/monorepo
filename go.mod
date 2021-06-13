module github.com/BradErz/monorepo

go 1.16

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	go.mongodb.org/mongo-driver v1.5.2
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.20.0
	go.opentelemetry.io/otel v0.20.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.20.0
	go.opentelemetry.io/otel/sdk v0.20.0
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
	google.golang.org/genproto v0.0.0-20210319143718-93e7006c17a6
	google.golang.org/grpc v1.37.1
	google.golang.org/protobuf v1.26.0
)
