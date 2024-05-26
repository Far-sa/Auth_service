# Stage 1: Build the application
FROM golang:1.22-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go mod download && go mod verify

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o /app/auth-svc ./cmd

# Stage 2: Create a minimal runtime image
#FROM alpine:latest
FROM alpine:latest

# Install necessary dependencies and remove apk cache in one step
RUN apk --no-cache add ca-certificates && rm -rf /var/cache/apk/*

# Copy the built Go app from the builder stage
COPY --from=builder /app/auth-svc /app/auth-svc

# COPY scripts/init-db.sh /docker-entrypoint-initdb.d/
# COPY ../../.env /docker-entrypoint-initdb.d/.env

# RUN chmod +x /docker-entrypoint-initdb.d/init-db.sh

# Expose port 50052 to the outside world
EXPOSE 50051

# Run the executable
CMD ["/app/auth-svc"]