FROM golang:1.13 AS builder
WORKDIR /go/src
COPY * .
RUN go build -o /go/bin/collector collector.go
CMD ["/go/bin/collector"]
