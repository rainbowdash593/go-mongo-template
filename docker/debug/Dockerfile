# Step 1: Modules caching
FROM golang:1.21.4-alpine3.17 as modules
COPY ../../go.mod ../../go.sum /modules/
WORKDIR /modules
RUN go mod download
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Step 2: Builder
FROM golang:1.21.4-alpine3.17 as builder
COPY --from=modules /go/pkg /go/pkg
COPY --from=modules /go/bin/dlv /go/bin/dlv
COPY ../.. /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} \
    go build -gcflags="all=-N -l" -tags migrate -o /bin/main ./cmd/main.go

# Step 3: Final
FROM alpine
COPY --from=builder /app/config /app/config
COPY --from=builder /bin/main /app/main
COPY --from=builder /go/bin/dlv /app/dlv
WORKDIR /app
ENTRYPOINT ["/app/dlv", "--listen=:${DELVE_PORT}", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/app/main"]