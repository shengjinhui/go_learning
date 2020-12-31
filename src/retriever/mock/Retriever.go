package mock

type Retriever struct {
	Contents string
}

// 在go中只要实现了相应的同名方法就算实现了,并不用继承相应的接口
func (r Retriever) Get(url string) string  {
	return r.Contents
}