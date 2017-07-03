# socket_server
jsonデータをTCPSocketで送るためのサーバー

## コンパイル方法
```sh
go build main.go
```

## 使い方
実行ファイルと同階層に`./assets`フォルダーを作る
その中にjsonファイルを入れる
例)main.json
## 実行方法
### Goの環境がある場合
`go run main.go -l` 
### 実行ファイルのみの場合
`{実行ファイル名} -l`