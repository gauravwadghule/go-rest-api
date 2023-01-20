FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Expose the port the application will run on
EXPOSE 5000

# Run the application
CMD ["./main"]
