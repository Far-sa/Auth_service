# Stage 1: Build the application
FROM golang:1.22-alpine AS builder

# Install dependencies
RUN apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o /app/auth-service

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

# Copy the built Go app from the builder stage
COPY --from=builder /app/auth-service .

# Expose port 50052 to the outside world
EXPOSE 50051

# Run the executable
CMD ["./auth-service"]