FROM golang

RUN mkdir /tmp/client
WORKDIR /tmp/client

ADD go.mod go.sum ./
RUN go mod download
ADD client.go ./
RUN go build -tags gocql_debug client.go

CMD ["./client"]
