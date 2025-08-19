# syntax=docker/dockerfile:1.7

FROM golang:1.24-alpine AS builder
WORKDIR /src

RUN apk add --no-cache ca-certificates tzdata git

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

ARG APP_PATH=./cmd/main.go
ARG BIN_NAME=auth


ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -trimpath -ldflags="-s -w" -o /out/${BIN_NAME} ${APP_PATH}


FROM gcr.io/distroless/static:nonroot
ENV GIN_MODE=release TZ=UTC \
    HTTP_PORT=8081 GRPC_PORT=50052 \
    GOMEMLIMIT=300MiB GOGC=75 GODEBUG=madvdontneed=1

WORKDIR /app
COPY --from=builder /out/auth /app/auth
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER nonroot:nonroot
EXPOSE 8081 50052
ENTRYPOINT ["/app/auth"]
