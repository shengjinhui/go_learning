package main

import (
	"fmt"
	"go_learning/src/oop/queue"
)

func main() {
	q := queue.Queue{}
	q.Push(1)
	q.Push(2)
	q.Push(4)
	q.Push(5)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
