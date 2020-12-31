package main

import "fmt"

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Printf("error occurred:%s", err)
		} else {
			panic(r)
		}
	}()
	b := 0
	a := 5 / b
	fmt.Println(a)
}
func main() {
	tryRecover()
}

