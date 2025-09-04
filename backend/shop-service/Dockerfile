# Stage 1: Build the Go binary
FROM golang:1.24.2

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o shop-service ./cmd/main.go

# Stage 2: Create a minimal final image
FROM debian:bullseye-slim

# Install necessary certificates
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy the built binary from builder
COPY --from=builder /app/shop-service .

# Optional: Set environment variables
ENV PORT=8084 \
    ENVIRONMENT=production

# Expose the port
EXPOSE $PORT

# Add a basic healthcheck (optional)
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD curl --fail http://localhost:$PORT/health || exit 1

CMD ["./shop-service"]
