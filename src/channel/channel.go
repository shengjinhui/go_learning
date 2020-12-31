package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for true {
		fmt.Printf("worker %d received %c\n", id, <-c)
	}
}

func chanDemo() {

	var chanels [10]chan int
	for i := 0; i < 10; i++ {
		chanels[i] = make(chan int)
		go worker(i, chanels[i])
	}
	for i := 0; i < 10; i++ {
		chanels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
