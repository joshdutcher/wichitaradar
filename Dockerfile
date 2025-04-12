# Build stage
FROM golang:1.22.2-alpine AS builder

WORKDIR /app

# Install modd for development
RUN go install github.com/cortesi/modd/cmd/modd@latest

# Copy source code
COPY . .

# Download dependencies and build the application
RUN go mod download && \
    go build -o wichitaradar cmd/server/main.go

# Development stage
FROM golang:1.22.2-alpine

WORKDIR /app

# Install modd
RUN go install github.com/cortesi/modd/cmd/modd@latest

# Copy the binary from builder
COPY --from=builder /app/wichitaradar .

# Copy templates and static files
COPY templates/ ./templates/
COPY static/ ./static/

# Expose port
EXPOSE 80

# Run the application
CMD ["./wichitaradar"]