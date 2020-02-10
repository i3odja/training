FROM golang:1.13.7-alpine3.11 AS builder

COPY . /go/

ENV GOPATH=/tmp
RUN go build -o main .

FROM alpine:3.11

COPY --from=builder /go/main /bin/
EXPOSE 8099

ENTRYPOINT ["/bin/main"]