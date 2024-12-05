
# Use a specific Go version with support for log/slog
FROM golang:1.21-alpine

# Set the working directory
WORKDIR /app

# Copy the Go source code
COPY . .

# Build the Go application

# RUN rm -rf /app/logs
# RUN mkdir /app/logs
# RUN touch /app/logs/app.log

RUN go mod tidy && go build -o auditapp .

# Set the entry point to run the application
CMD ["/app/auditapp"]

