package main

import (
	"path/filepath"
	"reflect"
	"runtime"
	"syscall/js"
)

var beforeUnloadCh = make(chan struct{})
var globalObj string = "GoWASM"

func beforeUnload(event js.Value) {
	beforeUnloadCh <- struct{}{}
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func registerCallback(fn func(args []js.Value)) js.Callback {
	dotFunctionName := filepath.Ext(GetFunctionName(fn))
	funcName := dotFunctionName[1:len(dotFunctionName)]

	callback := js.NewCallback(fn)

	wasmGo := js.Global().Get(globalObj)
	wasmGo.Set(funcName, callback)

	return callback
}

func blankCallback(args []js.Value) {

}

func cleanup() (js.Callback, chan struct{}) {
	beforeUnloadCb := js.NewEventCallback(0, beforeUnload)
	addEventListener := js.Global().Get("addEventListener")
	addEventListener.Invoke("beforeunload", beforeUnloadCb)

	return beforeUnloadCb, beforeUnloadCh
}

func callWASMLoad() js.Callback {

	callback := js.NewCallback(blankCallback)

	js.Global().Get("_onGoWASMLoad").Invoke(callback)

	return callback
}
