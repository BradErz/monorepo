buf-mod-update:
	buf mod update proto

buf: buf-mod-update
	buf format -w .
	buf generate proto/

format:
	gofumpt -w .
	buf format -w .

lint:
	golangci-lint run ./...

build:
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 \
		docker compose up --build --remove-orphans --force-recreate

install-tools:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
	go install mvdan.cc/gofumpt@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

grpcui:
	buf build --as-file-descriptor-set -o ./proto.bin
	grpcui -protoset ./proto.bin -plaintext localhost:10000