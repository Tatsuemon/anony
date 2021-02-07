# ANONY
短縮(匿名)URLサービス


やること

(CLI)

rpcに処理のerror内容を渡す

- Error モデルを作成して, Error Messageを全て格納

- [x] ユーザーの登録 + CLIの追加
    `$ anony register`
    で, name, email, password(confirm password)を入力させる

    name, emailはunique, passwordとconfirm passwordが異なる場合は, passwordだけ入力しなおす

    検証して, 登録 + ログイン

    ~トークンの発行 + ~/.anony/configに追加~
    JWTはsqliteに保存するようにする

- [ ] ユーザーの更新

- [ ] ユーザーの削除

- [x] ユーザーのログイン + CLIの追加
    `$ anony login`
    name(or email) passwordでログイン処理
    ~トークンの発行 + ~/.anony/configに追加~
    JWTはsqliteに保存するようにする

- [x] ユーザーの認証処理
    ユーザー登録, ログイン以外で使用するtokenの認証機能

- [ ] ユーザーのログアウト + CLIの追加
    `$ anony logout`

    Client側でJWTを削除

- [x] 登録APIの作成
    `$ anony create [original url]`
    original urlをshort URLに変更
    DBに追加

    一度stopになったURLも作り直せる

- [x] 自分のURLの一覧取得　status表示も
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
$ goose create CreateSequence sql
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

### sql 
COLLATE(称号順序)
utf8mb4_bin: UTF8の文字コード順(utf8mb4), バイナリコード順で全て区別する(bin)


### 値レシーバとポインタレシーバ

https://skatsuta.github.io/2015/12/29/value-receiver-pointer-receiver/


handlerでgrpcのmethodを使用する


### evansでデバッグ

https://narinymous.hatenablog.com/entry/2018/04/14/043908


### JWT

JWTの使い所
認証によく使われます。

クライアントは認証サーバに認証情報を渡し、トークンを請求する
認証サーバは認証情報が正しいことを確認して、秘密鍵を使って署名された JWT (user_id と expiration_date を含む) を発行する
クライアントはこの JWT を使って通常の API リクエストを行なう
サーバは、秘密鍵を使って JWT を検証し、user_id を JSON から取り出し、処理を行なう
JWTは改竄されていないことを確認できるため、改竄されていなければ認証サーバが署名したということなので、信用できる情報 (user_id) というわけです。

JWTのユースケース
OAuth 2.0のアクセストークン
OpenID ConnectのIDトークン
サービスをまたいだ認証機構
メール認証のトークン
JWTの構造
JWTの文字列の中には、ピリオド.が2つ含まれている。そこを境目にして3つの文字列に分割できる
1つ目: ヘッダ。どんなアルゴリズムで署名されているのか等のメタ情報を含む
2つ目: ペイロード。JSON本体に相当する情報
3つ目: 署名情報


https://qiita.com/arenahito/items/d96e437e5e13ef800ee0


CQRS
https://qiita.com/sanoyo/items/2cf5a8cae5928f37000d