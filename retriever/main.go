package main

import (
	"fmt"
	"godemo1/retriever/mock"
	"godemo1/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.baidu.com")
}
func main() {

	var r Retriever

	r = mock.Retriever{"This is a fake baidu.com"}
	fmt.Println(download(r))
	fmt.Printf("%T %v\n", r, r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)

	//	r = &test.Rr{"this is a test"}
	//	fmt.Printf("%T %v\n", r, r)
	//	fmt.Println(download(r))

	//fmt.Println(download(r))
}
