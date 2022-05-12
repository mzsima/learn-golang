# メモ
APIから取得したローソク足をWebページに表示する。

## フロント
- フロントはWASM syscall/jsを使ってみる

# compile WebAssembly
```
cd pages
GOOS=js GOARCH=wasm go build -o assets/main.wasm main.go
cp $(go env GOROOT)/misc/wasm/wasm_exec.js assets/
```


## 起動
```
go run server.go
```

## API
BTCのローソク足使いたいから、GMOコインの public api を使って取得　
https://api.coin.z.com/docs/#klines


