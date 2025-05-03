# Use the official Golang image as the base image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o url-shortener

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./url-shortener"]