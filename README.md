# shortURL
短縮URLサービス


やること

- [ ] 登録APIの作成

(input)
- url
(output)
- short_url

- [ ] locationを指定したpage(GET)の作成


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

## DBテスト
### テストデータの準備
Go製のテストデータの準備ツールを使っています。([レポジトリ](https://github.com/go-testfixtures/testfixtures))  
`db/fixtures`配下にテストデータをymlの形で入れています。  
`package interactor`で`prepareTestDB()`を呼ぶ事でテストデータがINSERTされます。  
※ testfixtureでは、テストデータの準備の前にDBをTRUNCATEするので、ローカルDBに対して使うのは要注意です。  
