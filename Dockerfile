# Stage 1: Build
FROM golang:1.24-alpine AS builder

# Environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the project
COPY . .

# Build the binary
RUN go build -o questionnaire

# Stage 2: Run
FROM alpine:latest

# Install CA certificates for HTTPS
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/questionnaire .

# Copy config folder if needed
COPY --from=builder /app/config ./config
COPY --from=builder /app/casbin ./casbin

# Expose default Gin port
EXPOSE 8080

# Set environment variable for production
ENV APP_ENV=production

# Run the app
CMD ["./questionnaire"]
