package main

import "fmt"

func main() {
	var a int = 10
	var p *int // 指针变量指向谁,就把谁的地址赋给指针变量
	p = &a
	fmt.Printf("p=%v, &a = %v\n", p, &a)

	*p = 666 // *p操作的不是p的内存,是p所指向的内存(就是a)
	fmt.Printf("*p = %v, a=%v\n", *p, a)
}
