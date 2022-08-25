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


# migration類
## migrationファイル作成手順
1. migrate create -ext sql -dir go/src/migration -seq {操作の名前(create_userとか、ファイル名に使われる)}
1. up.sqlに更新のSQL、down.sqlに切り戻しのSQLを書く

## migration更新類
基本的にgoを起動したときよしなにやってくれる、必要な時はdockerの中に入って作業を行う
- dockerの中に入るコマンド  
  `docker-compose exec go sh`
- migrationをリセットしたいとき  
  `migrate -path=migration -database="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@{host名(dockerの場合DB_CONTAINERのサービス名)}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" drop`
- migrationを手動アップデート  
  `migrate -path=migration -database="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@{host名(dockerの場合DB_CONTAINERのサービス名)}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" up {version}`
- migrationを手動ダウングレード  
  `migrate -path=migration -database="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@{host名(dockerの場合DB_CONTAINERのサービス名)}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" down {version}`


# 多分もうやらなくていい手順(作る際に利用した)
1. `docker-compose run --rm go /bin/ash`
1. `go mod init main`
1. `go mod download`
1. `go mod tidy`
1. `exit`
1. `docker-compose run --rm go go get github.com/golang-migrate/migrate/v4/database/postgres`
