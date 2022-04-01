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
		docker-compose up --build --remove-orphans --force-recreate

install-tools:
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	go install mvdan.cc/gofumpt@latest
