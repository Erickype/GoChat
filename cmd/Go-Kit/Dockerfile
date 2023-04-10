# Choose a base image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the source code to the container
COPY .. .

# Build the binary
RUN go build -o app ./main.go

# Start the service
CMD ["./app"]