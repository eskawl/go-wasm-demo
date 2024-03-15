package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("WASM Application!")

	js.Global().Call("updateDOM", "Set by WASM Application")
}

// we need to add the: "export add" comment above the function
//
//export add
func add(x int, y int) int {
	return x + y
}
