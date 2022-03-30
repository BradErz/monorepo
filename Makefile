buf:
	buf generate proto/

submodules:
	git submodule update --init --recursive

build:
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 \
		docker compose up --build

install-tools:
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	go install mvdan.cc/gofumpt@latest
