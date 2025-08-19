FROM golang:1.24-alpine AS builder
WORKDIR /src
RUN apk add --no-cache ca-certificates tzdata git build-base

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ARG APP_PATH=./cmd/main.go
ARG BIN_NAME=auth

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -trimpath -ldflags="-s -w -extldflags '-static'" -o /out/${BIN_NAME} ${APP_PATH}

FROM gcr.io/distroless/static:nonroot
ENV HTTP_PORT=8081 GRPC_PORT=50052 GIN_MODE=release TZ=UTC
WORKDIR /app
COPY --from=builder /out/auth /app/auth
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER nonroot:nonroot
EXPOSE 8081 50052
ENTRYPOINT ["/app/auth"]
