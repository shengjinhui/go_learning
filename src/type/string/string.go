package main

import "fmt"

func main() {
	var str string
	str = "hello world"
	ch := str[0]
	fmt.Printf(" the length of %s is %d\n", str, len(str))
	fmt.Printf(" the first character of %s is %c\n", str, ch)

	//str[0] = 'x' 编译错误,不允许修改初始化好的string
}
