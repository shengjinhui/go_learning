package main

type Integer int

// 面向对象
func (a Integer) Less(b Integer) bool {
	return a < b
}

// 面向过程
func Integer_Less(a Integer, b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

func (a Integer) Add2(b Integer) {
	a += b
}

func main() {
	var a Integer = 1
	println(a.Less(2))
	println(Integer_Less(a, 2))
	a.Add(2)
	a.Add2(2)
	println(a)
}
