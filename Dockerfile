FROM golang:1.21-alpine

WORKDIR /app

COPY echo_server.go .

RUN go build -o echo-server echo_server.go

EXPOSE 8080

CMD ["./echo-server"]

