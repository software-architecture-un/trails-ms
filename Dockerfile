FROM golang:latest

WORKDIR $GOPATH/src/trails-ms
COPY . .

RUN go get -d -v ./...
RUN go build

CMD ["./trails-ms"] 