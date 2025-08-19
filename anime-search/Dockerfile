# syntax=docker/dockerfile:1.7

FROM golang:1.24-alpine AS builder
WORKDIR /src

RUN apk add --no-cache ca-certificates tzdata git

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

ARG APP_PATH=./cmd/main.go
ARG BIN_NAME=anime-search

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -trimpath -ldflags="-s -w" -o /out/${BIN_NAME} ${APP_PATH}

FROM gcr.io/distroless/static:nonroot

ENV GIN_MODE=release \
    TZ=UTC \
    HTTP_PORT=8082 \
    GRPC_PORT=50053 \
    GOMEMLIMIT=300MiB \
    GOGC=75 \
    GODEBUG=madvdontneed=1

WORKDIR /app

COPY --from=builder /out/anime-search /app/anime-search
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER nonroot:nonroot
EXPOSE 8082 50053
ENTRYPOINT ["/app/anime-search"]
