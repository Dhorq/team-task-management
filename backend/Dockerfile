# backend/Dockerfile
FROM golang:1.26-alpine

WORKDIR /app

# Copy dependency files dulu
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Build binary
RUN go build -o main ./cmd/main.go

# Jalankan binary
CMD ["./main"]