package main

import (
	"syscall/js"
)

func getElementById(id string) js.Value {
	return js.Global().Get("document").Call("getElementById", id)
}

func addEventListener(element js.Value, event string, callback func(args []js.Value)) {
	element.Call("addEventListener", event, callback)
}
