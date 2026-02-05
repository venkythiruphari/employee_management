# Use the official Golang image as the base image
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN go build -o main ./cmd/main.go

# --- Start a new stage for a smaller, final image ---
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Copy the configuration file
COPY config.yaml .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
