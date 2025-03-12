FROM golang:1.24.1-alpine3.21 AS builder

RUN apk add --no-cache \
    git 

WORKDIR /app

ADD go.mod go.sum /app/
RUN go mod download

COPY . /app

WORKDIR /app/src
RUN go build \
    -ldflags "-X app/version.GitCommit=`git rev-parse --short=8 HEAD`" \
    -o /build/app

FROM alpine:3.21

RUN apk add --no-cache \
    make

WORKDIR /

COPY --from=builder /build/app /go-http-server-template
COPY makefile .

ENTRYPOINT [ "/go-http-server-template" ]
