# multistage docker build. This redices the size of the final docker image.
# stage 1 to build the app
FROM golang:alpine as builder

# Install CA certificates
RUN apk update && apk add --no-cache ca-certificates

# Set the GOPROXY environment variable
ENV GOPROXY=https://proxy.golang.org,direct


RUN mkdir /build 

ADD . /build/

WORKDIR /build 

RUN go build -o main .

# stage 2 deploys the app built in stage 1
FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3000

CMD ["./main"]