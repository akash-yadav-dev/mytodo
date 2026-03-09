# Build stage
FROM golang:alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make protobuf

WORKDIR /build

# Copy go mod files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Install protobuf codegen tools
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.11
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.6.1

# Regenerate protobuf files to avoid corrupted descriptors
RUN protoc -I apps/api/internal/auth/interfaces/grpc/proto \
	--go_out=apps/api/internal/auth/interfaces/grpc/pb --go_opt=paths=source_relative \
	--go-grpc_out=apps/api/internal/auth/interfaces/grpc/pb --go-grpc_opt=paths=source_relative \
	apps/api/internal/auth/interfaces/grpc/proto/auth.proto

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./apps/api/cmd/server

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./server"]
