FROM golang:1.19.0-alpine3.16

ENV GOPATH=

RUN apk update \
    && apk add git
    
WORKDIR /go/src
COPY ./src /go/src

EXPOSE 8080