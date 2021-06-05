package main

import (
	"syscall/js"
)

func moveLanding(this js.Value, arg []js.Value) interface{} {
	js.Global().Get("window").Get("location").Set(
		"href", "/",
	)
	return nil
}

func menuControl(this js.Value, arg []js.Value) interface{} {
	menu := js.Global().Get("document").Call(
		"getElementById", "header-slave",
	)
	menu.Get("classList").Call(
		"toggle", "visible",
	)
	return nil
}

func pageChanger(this js.Value, arg []js.Value) interface{} {
	js.Global().Get("window").Get("location").Set(
		"href", arg[0],
	)
	return nil
}

func main() {
	js.Global().Set("moveLanding", js.FuncOf(moveLanding))
	js.Global().Set("menuControl", js.FuncOf(menuControl))
	js.Global().Set("pageChanger", js.FuncOf(pageChanger))

	println("miho: wasm initiated")
	<-make(chan struct{})
	select {}
}
