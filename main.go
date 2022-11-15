package main

import (
	"syscall/js"
	"time"
)

func now() string {
	return time.Now().Format(time.RFC3339)
}

func main() {
	document := js.Global().Get("document")

	timestamp := document.Call("getElementById", "now")
	timestamp.Set("innerText", now())
}
