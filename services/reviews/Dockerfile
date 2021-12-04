FROM golang:1.17-alpine as base

WORKDIR /app

ENV CGO_ENABLED=0
COPY go.* ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . ./

FROM base AS build

ARG TARGETOS
ARG TARGETARCH
ARG PROJECT
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=linux GOARCH=amd64 go build -o server ${PROJECT}


FROM scratch AS release

# Copy the binary to the production image from the builder stage.
COPY --from=build /app/server /server
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Run the web service on container startup.
CMD ["/server"]