# Start from the official golang image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Command to run the executable
CMD ["./main"]
