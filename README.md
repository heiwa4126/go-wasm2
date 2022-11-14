# go-wasm2

Go(1.19)でWebAssemblyのサンプル。
ほぼ
[Interact with the DOM using Go's \`syscall/js\` module | egghead.io](https://egghead.io/lessons/go-interact-with-the-dom-using-go-s-syscall-js-module)
に基づく。

## 初期手順

```powershell
mkdir wasm2 ; cd wasm2
go mod init wasm1
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```
`wasm_exec.js`はGoのバージョンごとに変わるらしい。

あと、**このレポジトリには `wasm_exec.js`は入れてないので
上のcpはcloneしたところで実行してください。**

これに main.go と index.html を追加したのがこのレポジトリ。

## コンパイル

WindowsでPowerShellなら
```powershell
& { $Env:GOOS="js" ; $Env:GOARCH="wasm" ; go build -o main.wasm }
```

bashなら
```
GOOS=js GOARCH=wasm go build -o main.wasm
```

## テスト

で、Node.jsあるなら `npx http-server .`

http://127.0.0.1:8080/ 的なものが表示されるので、
これをブラウザで開いて(vscodeだと楽)、DOMが書き変わるのを見る。


## TinyGo

[TinyGo](https://tinygo.org/getting-started/install/)があるなら、
main.wasmがどれだけ小さくなるかが確認できる。

```bash
cp "$(tinygo env TINYGOROOT)/targets/wasm_exec.js" .
```
でTinyGo版のwasm_exec.jsをコピーして (std版とかなり違う)

```bash
tinygo build -target wasm -o main.wasm
# または
tinygo build -target wasm -o main.wasm -no-debug
```

# TODO

このままだと変化がないので、タイムスタンプとか表示してみる。
