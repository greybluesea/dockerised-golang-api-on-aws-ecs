FROM golang:latest

RUN mkdir -p /go/src/helloworld-api

WORKDIR /go/src/helloworld-api

COPY . /go/src/helloworld-api

# Build the Go application
RUN go build -o helloworld-api

# Set execute permissions on the binary
RUN chmod +x helloworld-api

CMD ["./helloworld-api"]

EXPOSE 80