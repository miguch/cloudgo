FROM golang:1.11

WORKDIR /go/src/github.com/miguch/cloudgo

COPY . .

RUN go get -d -v .
RUN go build -v .

