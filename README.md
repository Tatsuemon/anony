# ANONY
短縮(匿名)URLサービス

### インストール
```
$ make install-goose
```
### マイグレーションファイルの作成
```
$ goose create CreateSequence sql
```

### マイグレーションの実行
```
$ make goose-up-dev
```
### ロールバックの実行
```
$ make goose-down-dev
```

## テストの実行
```
$ make test
```

## テストカバレッジの確認
```
$ make coverage
```

## protobufの作成
```
$ make gen-proto
```