FROM golang:1.6.1-alpine

WORKDIR $GOPATH/src/github.com/rhperera/marvel-comic-api

COPY . .

RUN go get

EXPOSE 8080
