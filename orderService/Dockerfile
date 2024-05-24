# Build stage
FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /order-service cmd/main.go

# Final stage
FROM scratch

COPY --from=builder /order-service /order-service

EXPOSE 50055

ENTRYPOINT ["/order-service"]
