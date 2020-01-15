FROM golang:1.13 AS builder
WORKDIR /go/src
RUN pwd
COPY dummylogger.go .
RUN ls -l
RUN go build -o /go/bin/dummylogger dummylogger.go
RUN ls -l

FROM alpine:3
COPY --from=builder /go/bin/dummylogger .
CMD ["./dummylogger"]
