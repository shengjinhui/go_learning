package main

import (
	"fmt"
	"go_learning/src/retriever/mock"
	"go_learning/src/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) {
	poster.Post("http://www.baidu.com", map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
	// 也可以定义其他的方法
	Connect(host string)
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{Contents: "this is a fake baidu.com"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
