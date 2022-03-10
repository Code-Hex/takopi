package main

import (
	"syscall/js"

	"github.com/Code-Hex/takopi"
)

func takopilize(_ js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	return takopi.Do(args[0].String())
}

func registerCallbacks() {
	js.Global().Set("takopilize", js.FuncOf(takopilize))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	println("takopi WebAssembly Ready")
	<-c
}
