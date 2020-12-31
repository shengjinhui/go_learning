package main

import "time"

func main() {
	ch := make(chan int, 1024)
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()

	select {
	// 从ch中读取到数据
	case <-ch:
	// 一直没有从ch中读到数据,但从timeout中读取到数据
	case <-timeout:
	}
}
