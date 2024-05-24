# Stage 1: Build the application
FROM golang:1.18-alpine AS builder

# Install dependencies
RUN apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o /product-catalog-service cmd/main.go

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the built Go app from the builder stage
COPY --from=builder /product-catalog-service .

# Expose port 50053 to the outside world
EXPOSE 50053

# Run the executable
CMD ["./product-catalog-service"]
