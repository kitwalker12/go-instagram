FROM golang

COPY . /go/src/github.com/hieven/go-instagram
WORKDIR /go/src/github.com/hieven/go-instagram

RUN go get ./...
RUN go build
