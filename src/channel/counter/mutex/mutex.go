package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 使用锁来实现counter递增
var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func main() {
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go Count(lock)
	}

	// 这里使用for循环不断检查counter的值,大于10,则程序结束,退出主函数
	for {
		lock.Lock()
		c := counter
		lock.Unlock()

		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}
