FROM golang:1.19.0-alpine3.16

ENV GOPATH=

RUN apk update \
    && apk add git \
    && go mod init main \
    && go mod download \
    && go mod tidy 

WORKDIR /go/src
COPY ./src /go/src

EXPOSE 8080