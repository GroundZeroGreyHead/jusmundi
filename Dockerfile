FROM golang:1.20-alpine

# Install Go dependencies
RUN apk add --no-cache build-base ca-certificates git

# Set the working directory
WORKDIR /app

# Copy the application code
COPY go.mod go.sum main.go ./

# Build the application
RUN go build -o main main.go

# Command to run the application
CMD ["/app/main"]