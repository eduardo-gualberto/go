# Development Dockerfile for Go API (go1.24.2)
#
# Build an interactive shell with your current code mounted:
#
#   docker build -t go-api-dev .
#   docker run -it --rm -v $(pwd):/app -w /app -p 8080:8080 go-api-dev

FROM golang:1.24.2

# Enable Go modules
ENV GO111MODULE=on

# Set working directory inside the container
WORKDIR /app

# Cache go.mod downloads
COPY go.mod go.sum ./
RUN go mod download

# Expose the application's HTTP port
EXPOSE 8080

# Default to bash for interactive development
CMD ["go", "run", "cmd/dummy/main.go"]