# Stage 1: Build
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 👇 Если main.go в cmd/server, оставьте так
RUN go build -o main ./cmd/server

# Stage 2: Run
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./main"]