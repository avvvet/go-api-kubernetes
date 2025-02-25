# Use an official Golang image to build the app
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy just the Go module files to cache dependencies
COPY go.mod ./

# Download and install Go module dependencies
RUN go mod download

# Copy the entire application directory into the container
COPY . .

# Build the Go application
RUN go build -o api-service .

# Expose the port the app runs on
EXPOSE 8080

# Run the Go app
CMD ["./api-service"]
