FROM golang:1.22-alpine

WORKDIR /app

# Install required system packages
RUN apk add --no-cache gcc musl-dev

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd/server

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./main"]