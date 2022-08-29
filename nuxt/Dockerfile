FROM node:lts-alpine

RUN mkdir ./src
# コマンド実行
RUN apk update && \
    apk upgrade && \
    apk add git && \
    npm install -g npm@latest nuxi nuxt3

EXPOSE 3000
ENV HOST 0.0.0.0