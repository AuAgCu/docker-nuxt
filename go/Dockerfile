FROM golang:1.19.0-alpine3.16

# ENV GOPATH=

RUN apk update \
    && apk add git \
    && apk add curl \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz | tar xvz \
    && mv ./migrate.linux-amd64 /usr/bin/migrate \
    && go install github.com/cosmtrek/air@latest
    
WORKDIR /go/src
# COPY ./src /go/src

EXPOSE 3001