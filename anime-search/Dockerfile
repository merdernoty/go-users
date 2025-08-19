FROM golang:1.24-alpine AS builder
WORKDIR /src
RUN apk add --no-cache ca-certificates tzdata git build-base

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ARG APP_PATH=./cmd/main.go
ARG BIN_NAME=anime-search

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -trimpath -ldflags="-s -w -extldflags '-static'" -o /out/${BIN_NAME} ${APP_PATH}

FROM gcr.io/distroless/static:nonroot
ENV HTTP_PORT=8082 GRPC_PORT=50053 GIN_MODE=release TZ=UTC
WORKDIR /app
COPY --from=builder /out/anime-search /app/anime-search
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER nonroot:nonroot
EXPOSE 8082 50053
ENTRYPOINT ["/app/anime-search"]
