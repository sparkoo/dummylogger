FROM golang:1.13 AS builder
WORKDIR /go/src
COPY collector/* .
RUN go build -o /go/bin/collector collector/collector.go

FROM alpine:3
COPY --from=builder /go/bin/dummylogger .
CMD ["./collector"]
