package test

import (
	"fmt"
)

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Print(a, b, s)
}

func main() {
	fmt.Print("Hello World")
	variableInitialValue()
}
