# Stage 1: Build the application
FROM golang:1.20-alpine AS builder

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
RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service cmd/main.go

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the built Go app from the builder stage
COPY --from=builder /app/auth-service .

# Expose port 50052 to the outside world
EXPOSE 50052

# Run the executable
CMD ["./auth-service"]