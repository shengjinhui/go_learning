package main

import "fmt"

func main() {

	// 三种初始化数组的方式
	var arr [3]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 45, 5}

	var grid [4][5]int

	fmt.Println(arr, arr2, arr3)
	fmt.Println(grid)

	for i := range arr2 {
		fmt.Println(i, arr2[i])
	}

	for i, v := range arr2 {
		fmt.Println(i, v)
	}
}
