
FROM golang:1.21.3 as build
WORKDIR /src
COPY go.mod .
RUN go mod download

COPY *.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/helloworld-api 

FROM alpine:3.13
COPY --from=build /bin/helloworld-api /bin/helloworld-api
EXPOSE 8080
CMD ["/bin/helloworld-api"]