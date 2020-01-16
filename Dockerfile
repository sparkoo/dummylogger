FROM golang:1.13 AS builder
WORKDIR /go/src
COPY dummylogger.go .
RUN go build -o /go/bin/dummylogger dummylogger.go

FROM alpine:3
COPY --from=builder /go/bin/dummylogger .
CMD ["./dummylogger"]
