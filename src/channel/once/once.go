package main

import (
	"sync"
	"time"
)

var a string
var once sync.Once

func setUp() {
	a = "hello world"
}

func doprint() {
	once.Do(setUp)
	println(a)
}

func twoprint() {
	go doprint()
	go doprint()
}
func main() {
	twoprint()
	time.Sleep(1e9)
}
