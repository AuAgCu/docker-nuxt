FROM golang:1.19.0-alpine3.16

ENV GOPATH=

RUN apk update \
    && apk add git \
    && apk add curl \
    # && apk add install -y curl gnupg2 vim \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz | tar xvz \
    && mv ./migrate.linux-amd64 /usr/bin/migrate
    
WORKDIR /go/src
COPY ./src /go/src

EXPOSE 8080