package main

import (
	"fmt"
	"syscall/js"
)

func main() {

	defer registerCallback(goPrint).Release()
	defer registerCallback(higherOrderCallback).Release()

	// callback for when wasm is loaded
	wasmCB := callWASMLoad()
	defer wasmCB.Release()

	// cleanup for when the page unloads
	cleanupCb, cleanupCh := cleanup()
	defer cleanupCb.Release()
	<-cleanupCh
	fmt.Println("go-wasm exit 0")
}

func higherOrderCallback(args []js.Value) {
	isNice := args[0].Bool()
	setRender := js.Global().Get("setRender")

	renderText := render(isNice)
	setRender.Invoke(renderText)
}

func goPrint(args []js.Value) {
	message := args[0].String()
	fmt.Println(message)
}

func render(nice bool) string {
	if nice == true {
		return "hello"
	}
	return "bye"
}
