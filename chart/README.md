# Goでチャート遊び

## フロント
- フロントはWASM syscall/jsを使ってみる

# compile WebAssembly
```
GOOS=js GOARCH=wasm go build -o assets/main.wasm
cp $(go env GOROOT)/misc/wasm/wasm_exec.js assets/
```

