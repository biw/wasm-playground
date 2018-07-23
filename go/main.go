package main

import (
	"fmt"
	"syscall/js"
)

var done = make(chan struct{})

func main() {
	callback := js.NewCallback(higherOrderCallback)
	defer callback.Release()
	setPrintMessage := js.Global().Get("setPrintMessage")
	setPrintMessage.Invoke(callback)
	<-done
}

func higherOrderCallback(args []js.Value) {
	isNice := args[0].Bool()
	setRender := js.Global().Get("setRender")

	renderText := render(isNice)
	setRender.Invoke(renderText)
}

func printMessage(args []js.Value) {
	message := args[0].String()
	fmt.Println(message)
	done <- struct{}{}
}

func render(nice bool) string {
	if nice == true {
		return "dis is nice"
	}
	return "dis is mean"
}
