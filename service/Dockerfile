# Base image
FROM golang:latest

# Set working directory
WORKDIR /app

# Copy source files
COPY . .
RUN go mod download

# Build the app
RUN go build

# Expose port 8080
EXPOSE 8000

# Run the app
CMD ["./UrlShortner"]
