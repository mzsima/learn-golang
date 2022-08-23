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
