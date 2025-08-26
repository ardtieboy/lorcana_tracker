# --- Build Stage ---
# Use an official Go image to build the application
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies first
# This leverages Docker layer caching
COPY go.mod go.sum ./
RUN apk add --no-cache gcc musl-dev
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
# CGO_ENABLED=1 allows go-sqlite3 to work
# -o /app/main creates the output file in the /app directory
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o /app/main .

# --- Final Stage ---
# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Copy the database files
# Note: In a production scenario, you would typically use a managed database
# instead of including the database file in the container.
COPY lorcana.db .

# Expose the port the application runs on
EXPOSE 8080

# The command to run when the container starts
# We run the app and pass the --initDB flag to ensure the database is ready.
# You might adjust this depending on your production needs.
CMD ["/app/main", "--initDB"] 
