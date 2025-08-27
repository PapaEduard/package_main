# Use the official Golang image as a build stage
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to cache dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o metrics-app

# Use a minimal image for the final stage
FROM alpine:latest

# Install ca-certificates for HTTPS support
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the compiled binary from builder stage
COPY --from=builder /app/metrics-app .

# Expose the port the app listens on
EXPOSE 9200

# Command to run the executable
ENTRYPOINT ["./metrics-app"]
