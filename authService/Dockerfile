#!! gpt -generate
# Stage 1: Build the Go app with gRPC support
FROM golang:1.22-alpine AS builder

# Install necessary packages
RUN apk add --no-cache git make protobuf

# Install protoc-gen-go and protoc-gen-go-grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY authService/go.mod authService/go.sum ./

# Download Go modules
RUN go mod download

# Copy the entire project directory into the container
COPY authService/ .

# Copy the shared proto files from the root project directory
COPY shared-proto/ ./shared-proto/

# Copy the third-party proto files
COPY third_party/ ./third_party/

# Download the necessary Google protobuf files
RUN mkdir -p /usr/include/google/protobuf && \
    wget -qO- https://raw.githubusercontent.com/protocolbuffers/protobuf/main/src/google/protobuf/descriptor.proto > /usr/include/google/protobuf/descriptor.proto && \
    wget -qO- https://raw.githubusercontent.com/protocolbuffers/protobuf/main/src/google/protobuf/timestamp.proto > /usr/include/google/protobuf/timestamp.proto && \
    wget -qO- https://raw.githubusercontent.com/protocolbuffers/protobuf/main/src/google/protobuf/any.proto > /usr/include/google/protobuf/any.proto && \
    wget -qO- https://raw.githubusercontent.com/protocolbuffers/protobuf/main/src/google/protobuf/duration.proto > /usr/include/google/protobuf/duration.proto && \
    wget -qO- https://raw.githubusercontent.com/protocolbuffers/protobuf/main/src/google/protobuf/empty.proto > /usr/include/google/protobuf/empty.proto && \
    wget -qO- https://raw.githubusercontent.com/protocolbuffers/protobuf/main/src/google/protobuf/struct.proto > /usr/include/google/protobuf/struct.proto && \
    wget -qO- https://raw.githubusercontent.com/protocolbuffers/protobuf/main/src/google/protobuf/wrappers.proto > /usr/include/google/protobuf/wrappers.proto

# Generate protobuf files
RUN protoc -I=./shared-proto -I=./third_party -I=/usr/include \
    --go_out=./pb --go-grpc_out=./pb \
    ./shared-proto/auth/*.proto \
    ./shared-proto/authorize/*.proto

# Build the Go app
RUN go build -o /app/auth-service ./cmd/main.go

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Install necessary tools
RUN apk add --no-cache ca-certificates postgresql-client

# Set the current working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/auth-service /app/auth-service

# Expose port for gRPC
EXPOSE 50051

# Command to run the service
CMD ["/app/auth-service"]

#!!!
# FROM golang:1.22-alpine AS builder

# # Set the current working directory inside the container
# WORKDIR /app

# # Copy go mod and sum files
# COPY go.mod go.sum ./

# RUN go mod download

# # Copy the source code into the container
# COPY . .

# # Copy the shared proto files
# COPY ./shared-proto/*.proto ./pb/

# # Build the Go app
# RUN go build -o /app/auth-service ./cmd

# # Minimal runtime image
# FROM alpine:latest

# # Install necessary tools
# RUN apk add --no-cache ca-certificates postgresql-client

# # Copy the built binary from the builder
# COPY --from=builder /app/auth-service /app/auth-service

# # Expose port for your service
# EXPOSE 50051

# # Command to run your service
# CMD ["/app/auth-service"]

#!! multi stage builder for proto

# First stage: Build the proto files
# FROM golang:1.16 as protobuilder

# # Install protoc and the Go plugin for protoc
# RUN apt-get update && apt-get install -y protobuf-compiler && \
#     go get google.golang.org/protobuf/cmd/protoc-gen-go

# # Download the Google APIs
# RUN git clone https://github.com/googleapis/googleapis.git

# WORKDIR /proto
# COPY ./shared-proto/*.proto ./
# RUN protoc --proto_path=. --proto_path=googleapis --go_out=. *.proto

# # Second stage: Build the Go app
# FROM golang:1.22 as builder

# WORKDIR /app
# COPY . .
# COPY --from=protobuilder /proto/*.pb.go ./pb/

# RUN go mod download
# RUN go build -o /app/auth-service ./cmd

# # Final stage: Create the runtime image
# FROM alpine:latest

# RUN apk add --no-cache ca-certificates postgresql-client

# WORKDIR /app
# COPY --from=builder /app/auth-service .

# EXPOSE 50051
# CMD ["./auth-service"]
#!!! use with buf package
# First stage: Build the proto files
# FROM golang:1.16 as protobuilder

# # Install buf
# RUN apt-get update && apt-get install -y curl && \
#     curl -sSL https://github.com/bufbuild/buf/releases/download/v0.43.2/buf-Linux-x86_64 -o /usr/local/bin/buf && \
#     chmod +x /usr/local/bin/buf

# # Download the Google APIs
# RUN git clone https://github.com/googleapis/googleapis.git /proto/googleapis

# WORKDIR /
# COPY ./shared-proto/*.proto ./proto/
# COPY buf.yaml ./proto/

# WORKDIR /proto
# RUN buf generate --path . --path ./googleapis

# # Second stage: Build the Go app
# FROM golang:1.22 as builder

# WORKDIR /app
# COPY . .
# COPY --from=protobuilder /proto/*.pb.go ./pb/

# RUN go mod download
# RUN go build -o /app/auth-service ./cmd

# # Final stage: Create the runtime image
# FROM alpine:latest

# RUN apk add --no-cache ca-certificates postgresql-client

# WORKDIR /app
# COPY --from=builder /app/auth-service .

# EXPOSE 50051
# CMD ["./auth-service"]
#! important
# FROM alpine:latest AS builder

# # Set the current working directory inside the container
# WORKDIR /app

# # Copy go mod and sum files
# COPY go.mod go.sum ./

# RUN go mod download && go mod verify

# # Copy the source code into the container
# COPY . .

# # Builder stage doesn't build the binary (handled by Makefile)

# # Minimal runtime image
# FROM alpine:latest

# # Install necessary dependencies (e.g., database driver for your chosen language)
# # (Optional, if your service requires them)
# RUN apk add --no-cache <database_driver_package> && rm -rf /var/cache/apk/*

# # Copy the built binary from the host (assuming Makefile builds it)
# COPY ./bin/auth-service /app/auth-service  

# # Expose port for your service (replace with your service's port)
# EXPOSE 8080

# # Command to run your service
# CMD ["/app/auth-service"]
