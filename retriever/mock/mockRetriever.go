package mock

import "fmt"

type Retriever struct {
	Contents string
}

// 类似java tostring
func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever:{Contents=%s}", r.Contents)
}

// 值接收者
func (r *Retriever) Get(url string) string {
	return r.Contents
}

// 实现Poster接口
func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}
