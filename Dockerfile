# First stage of the build process
# Alpine is a lightweight linux distro
# golang:alpine combines the Go language with the Alpine distro
FROM golang:alpine as build
LABEL authors="ThomasCanning"

# Install git and bash in the Alpine image
# Needed for go mod download
RUN apk add --no-cache git bash

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY ./cmd/web .

# Build the Go app
RUN go build -o main .

# Final stage
# alpine:latest is a very small image, smaller than the golang:alpine image because that also contains the compiler
FROM alpine:latest

# Copy the compiled Go binary from the build stage
# Multi stage builds reduce image size because the final image only contains the compiled binary and not the source code
COPY --from=build /app/main /main

# Expose port 4000 to the outside world
EXPOSE 4000

# Command to run the executable
CMD ["./main"]
