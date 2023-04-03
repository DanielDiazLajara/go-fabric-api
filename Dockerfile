# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY *.go ./
COPY ./src ./src

# Build the Go app
RUN go build -o main ./

# Expose port 8080 to the outside world
EXPOSE 3000

# Run the executable
CMD ["./main"]