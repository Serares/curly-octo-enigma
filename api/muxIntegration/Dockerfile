# Dockerfile for app go module
# Use the official Golang image as a base image
FROM golang:1.20-alpine

# Install SQLite
RUN apk add --no-cache sqlite

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .
COPY ../../data/questions-app /app/questions-app
# Build the Go app
RUN go build -o main .

# Set environment variables for SQLite database path
ENV SQLITE_DB_PATH=/app/questions-app
# Command to run the executable
CMD ["./main"]
