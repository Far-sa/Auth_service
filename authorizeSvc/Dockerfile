# Use the official Golang image as the builder
FROM golang:1.22-alpine AS builder


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Copy the entire project
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Build the Go app
RUN go build -o /app/authz-svc ./cmd

# Start a new stage from scratch
FROM alpine:latest


# Copy the pre-built binary from the builder stage
COPY --from=builder /app/authz-svc /app/authz-svc

# COPY /scripts/init-db.sh /docker-entrypoint-initdb.d/
# COPY ../../.env /docker-entrypoint-initdb.d/.env

# RUN chmod +x /docker-entrypoint-initdb.d/init-db.sh
# Expose port 8080 to the outside world
EXPOSE 50052

# Command to run the executable
CMD ["/app/authz-svc"]


# # Stage 1: Build the application
# FROM golang:1.22-alpine AS builder

# # Set the current working directory inside the container
# WORKDIR /app/cmd

# # Copy go mod and sum files
# COPY go.mod go.sum ./
# RUN go mod download && go mod verify

# # Copy the source code into the container
# COPY . /app 
# RUN ls -la

# # Build the Go app
# RUN go build -o /app/authz-svc

# # Stage 2: Create a minimal runtime image
# FROM alpine:latest


# # Copy the built Go app from the builder stage
# COPY --from=builder /app/authz-svc .

# # Expose port 50052 to the outside world
# EXPOSE 50052

# # Run the executable
# CMD ["./authz-svc"]
