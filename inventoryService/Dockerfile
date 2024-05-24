# Build stage
FROM golang:1.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /inventory-service cmd/main.go

# Run stage
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /inventory-service /inventory-service

EXPOSE 50051

ENTRYPOINT ["/inventory-service"]
