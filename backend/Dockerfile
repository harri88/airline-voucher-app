# Stage 1: Builder
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache build-base

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Generate Wire
RUN go generate ./...

# Build the binary
# CGO_ENABLED=1 is required for go-sqlite3
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o /server ./cmd/server

# Stage 2: Final Image
FROM alpine:latest

RUN apk add --no-cache ca-certificates

# Set working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /server .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./server"]