FROM alpine:latest

# Install certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create the application directory
RUN mkdir /app

# Copy the pre-built binary to the /app directory
COPY content-fetcher /app/content-fetcher

# Copy the .env file to the /app directory
# COPY .env /app/.env

# Set the working directory
WORKDIR /app

# Set the command to run the binary
CMD ["/app/content-fetcher"]
