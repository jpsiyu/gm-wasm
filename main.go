// +build js

package main

import (
	"encoding/json"
	"syscall/js"
)

func register() {
	js.Global().Set("goFunParse", js.FuncOf(parse))
}

func parse(this js.Value, args []js.Value) interface{} {
	method := args[0].String()
	params := args[1].String()
	fee := int64(args[2].Int())

	response := CMD(method, params, fee)
	result, _ := json.Marshal(response)
	return string(result)
}

func main() {
	register()
	select {}
}
