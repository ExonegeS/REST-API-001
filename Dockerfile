# Use a minimal base image for Go binaries
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the pre-built binary from the local `bin/` directory to the container
COPY bin/app /app/app

# Set the port that the service listens on
EXPOSE 8888

# Set permissions to make the binary executable
RUN chmod +x /app/app

# Set the command to run the service
CMD ["./app"]