# 環境構築
1. https://www.docker.com/get-started/からアプリをダウンロード
1. ターミナルを起動して下記コマンド  
  `$ brew install docker`
1. docker-composeをインストール  
  `$ brew install docker-compose`
1. ビルド  
  `$ docker-compose build`
1. npm install  
  `$ docker-compose run --rm nuxt npm install`
1. 起動  
  `$ docker-compose up`
1. ブラウザからlocalhost:9001にアクセス  

# 多分もうやらなくていい手順(作る際に利用した)
1. `docker-compose run --rm go /bin/ash`
1. `go mod init main`
1. `go mod download`
1. `go mod tidy`
1. `exit`