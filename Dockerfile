# Step 1: Use an official Go image to build the Go app
FROM golang:1.24 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Step 2: Build the Go app
WORKDIR /app/cmd/server


RUN GOOS=linux GOARCH=amd64 go build -o main .

# Step 3: Use a smaller base image to run the app
FROM alpine:latest

# Install necessary dependencies for running the Go binary
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the built Go binary from the build stage
COPY --from=build /app/cmd/server/main .

COPY --from=build /app/cmd/server/.env /root/.env

# Expose the port that the Go app will listen on
EXPOSE 8090

# Command to run the Go binary
CMD ["./main"]