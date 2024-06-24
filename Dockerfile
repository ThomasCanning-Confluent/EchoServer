FROM golang:latest
LABEL authors="ThomasCanning"

# Add Maintainer Info
LABEL maintainer="ThomasCanning <tcanning@confluent.io>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the cmd/web directory to the Working Directory inside the container
COPY ./cmd/web .

# Build the Go app
RUN go build -o main .

# Expose port 4000 to the outside world
EXPOSE 4000

# Command to run the executable
CMD ["./main"]