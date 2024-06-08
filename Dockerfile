# Use an official Golang image as a base
FROM golang:alpine

# Set the working directory to /app
WORKDIR /app

# Copy the Go module files
COPY go.mod ./

# Install the dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go program
RUN go build -o main main.go

# Run the command to start the server
CMD ["go", "run", "main.go"]