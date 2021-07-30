FROM golang:1.16.6-alpine3.14 as builder
RUN apk update && apk upgrade && \
    apk add \
    xz-dev \
    musl-dev \
    gcc
RUN mkdir -p /go/src/github.com/alexsosic/fizz-buzz-api
COPY . /go/src/github.com/alexsosic/fizz-buzz-api
RUN cd /go/src/github.com/alexsosic/fizz-buzz-api && env CGO_ENABLED=1 go build

FROM alpine:3.14
RUN apk update && apk upgrade && \
        apk add --no-cache ca-certificates xz postgresql-client
COPY ./entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
COPY --from=builder /go/src/github.com/alexsosic/fizz-buzz-api/fizz-buzz-api /usr/bin
ENTRYPOINT ["/entrypoint.sh", "--", "/usr/bin/fizz-buzz-api"]

EXPOSE 8080
