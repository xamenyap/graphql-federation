# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first (leveraging cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o review main.go

# Stage 2: Run
FROM scratch

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/review .

# Expose port
EXPOSE 8082

# Run the binary
ENTRYPOINT ["./review"]
