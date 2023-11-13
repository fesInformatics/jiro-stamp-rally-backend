#  二郎スタンプラリー BE

## 概要

- golang v1.20
- gorm(予定) → 使わずに頑張りたい

## 環境構築

Dokcer, DockerComposeがインストールされている前提です。

```bash
cp .env.example .env
```

以下のコマンドで[Air](https://github.com/cosmtrek/air)が起動します。

```bash
docker compose up -d
```

`http://localhost:3000/health`をブラウザで開いて、`{"message":"OK"}`が表示されることを確認してください。

Airはファイルの変更検知してホットリロードします。

## マイグレーション方法
以下のコマンドでworkspaceコンテナに入って、コマンドを実行してください。

```
docker compose exec workspace sh
```

```sh
マイグレーション up
migrate -database="mysql://$DB_USER:$DB_PASSWORD@tcp(mysql:$DB_PORT)/$DB_DATABASE?multiStatements=true" -path=jiro-stamp-rally/db/migrations/ up 1
マイグレーション down
migrate -database="mysql://$DB_USER:$DB_PASSWORD@tcp(mysql:$DB_PORT)/$DB_DATABASE?multiStatements=true" -path=jiro-stamp-rally/db/migrations/ down 1
マイグレーション バージョン確認
migrate -database="mysql://$DB_USER:$DB_PASSWORD@tcp(mysql:$DB_PORT)/$DB_DATABASE?multiStatements=true" -path=jiro-stamp-rally/db/migrations/ version
マイグレーション force※
migrate -database="mysql://$DB_USER:$DB_PASSWORD@tcp(mysql:$DB_PORT)/$DB_DATABASE?multiStatements=true" -path=jiro-stamp-rally/db/migrations/ force 1
マイグレーションファイル作成
migrate create -ext sql -dir /jiro-stamp-rally/db/migrations -seq [マイレーションファイル名]
```
※マイグレーションの際にSQLの実行に失敗するとバージョンがdirtyになりマイグレーション操作が行えなくなる。forceをすると、強制的にdirtyを外すことができる。
