FROM golang:latest AS builder
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build cmd/main.go
CMD ["/app/main"]