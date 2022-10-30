FROM cgr.dev/chainguard/go:1.19 AS build

ARG TARGETOS
ARG TARGETARCH
ARG PROJECT

WORKDIR /app

ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg \
    go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    GOOS=linux GOARCH=amd64 go build -o server ${PROJECT}

# Final image
FROM cgr.dev/chainguard/static:latest

COPY --from=build /app/server /server

CMD ["/server"]