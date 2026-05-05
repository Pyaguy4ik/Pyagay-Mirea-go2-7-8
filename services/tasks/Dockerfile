# Stage 1: Build stage
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/bin/tasks ./cmd/tasks

# Stage 2: Runtime stage
FROM alpine:3.20

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/bin/tasks /app/tasks

# Expose port
EXPOSE 8082

# Run the application
CMD ["/app/tasks"]
