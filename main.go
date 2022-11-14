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
	// document = window.document;

	myH1 := document.Call("createElement", "h1")
	// myH1 = document.createElement("h1");

	myH1.Set("innerText", "This is H1")
	// myH1.innerHTML = "This is H1";

	document.Get("body").Call("appendChild", myH1)
	// document.body.appendChild(myH1);

	timestamp := document.Call("getElementById", "now")
	timestamp.Set("innerText", now())
}
