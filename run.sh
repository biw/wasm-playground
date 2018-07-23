GOOS=js GOARCH=wasm go1.11beta2 build -o main.wasm go/main.go go/dom.go go/lib.go && python3 server.py $1
