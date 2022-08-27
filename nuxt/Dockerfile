FROM node:16.13.2-alpine

EXPOSE 3000
ENV HOST 0.0.0.0

WORKDIR /nuxt/src

# コマンド実行
RUN apk update && \
    apk add git && \
    npm install && \
    npm install -g @vue/cli nuxt create-nuxt-app