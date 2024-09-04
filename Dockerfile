# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy the project files
COPY . .

# Change directory to the cmd/gin folder
WORKDIR /app/cmd/gin

# Download dependencies
RUN go mod tidy

# Build the binary
RUN go build -o /app/go-budget

# Stage 2: Run the application using a smaller base image
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/go-budget /app/go-budget

# Copy the templates directory from the current environment
COPY pkg/templates /app/pkg/templates

# Run the binary
CMD ["./go-budget"]