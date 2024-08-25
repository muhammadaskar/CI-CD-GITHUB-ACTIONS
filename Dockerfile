# Use golang:latest as the base image
FROM golang:latest

# Build argument
ARG APP_PORT
ENV APP_PORT=${APP_PORT}

# Set the working directory to /app-prod
WORKDIR /app-prod

# Copy the contents of the current directory to /app-prod in the container
COPY . .

# Install dependencies (if needed)
RUN go get -d -v ./...

# Compile the Golang application
RUN go build -o main .

# Command to run when the container starts
CMD ["./main"]
