# Use an official Go runtime as a parent image
FROM golang:1.17 AS build

# Set the working directory
WORKDIR /go/src/app

# Copy the Go application source code to the container
COPY app/ .

# Build the Go application
RUN go get -d -v 
RUN go install -v 
RUN go build -o /go/bin/app

# Use a minimal base image for the final executable
FROM gcr.io/distroless/base-debian10

# Set the working directory
WORKDIR /app

# Copy the built Go binary from the build stage
COPY --from=build /go/bin/app .

# Expose the port your Go application listens on
EXPOSE 8081

# Define the command to run your Go application
CMD ["./app"]
