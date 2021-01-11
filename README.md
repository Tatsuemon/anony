# ANONY
短縮(匿名)URLサービス


やること

(CLI)
- [ ] ユーザーの登録 + CLIの追加
    `$ anony register`
    で, name, email, password(confirm password)を入力させる

    name, emailはunique, passwordとconfirm passwordが異なる場合は, passwordだけ入力しなおす

    検証して, 登録 + ログイン
    トークンの発行 + ~/.anony/configに追加

- [ ] ユーザーのログイン + CLIの追加
    `$ anony login`
    name(or email) passwordでログイン処理
    トークンの発行 + ~/.anony/configに追加

- [ ] ユーザーのログアウト + CLIの追加
    `$ anony logout`

    トークンの削除, DBのAuths Tableのstatusの変更

- [ ] 登録APIの作成
    `$ anony create [original url]`
    original urlをshort URLに変更
    DBに追加

    一度stopになったURLも作り直せる

- [ ] 自分のURLの一覧取得　status表示も
    `$ anony list`

    optionで稼働しているもののみ取得

- [ ] shortURLのはいし
    `$ anony drop [-u original_url] [-i url_id]`
    statusの変更

(WEBサーバー)
- [ ] locationを指定したpage(GET)の作成

### DB設計

Users Table
id, name, password, email

name (or email), passwordでログイン

Auths Table
user_id, token, status(0: inactive, 1: active)

statusはそのトークンを使用できるかどうか

登録は, 全てを入力

CLIツールから利用する場合は, 最初のログイン完了後にトークンを~/.anony/configに追加する
認証機能はそのトークンを使用して, ユーザーの特定


URLs Table
id, original, short, status(0: inactive, 1: active), user_id


作成されるURLは
http://hoge.com/user_nameの変形/ランダム文字列


### インストール
`make install-goose`

### マイグレーションファイルの作成
```
$ goose -dir "db/migrations" create CreateSequence sql
```

### マイグレーションの実行
```
$ docker-compose run app goose up
```
### ロールバックの実行
```
$ docker-compose run app goose down
```

### grpc
参考: https://qiita.com/marnie_ms4/items/4582a1a0db363fe246f3