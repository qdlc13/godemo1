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
	//r不是简单的引用 有类型%T和值%v
	var r Retriever
	//r = &mock.Retriever{"This is a fake baidu.com"}也可以
	r = mock.Retriever{"This is a fake baidu.com"}
	fmt.Println(download(r))
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	//判断类型的等价写法
	//Type Assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut) //1m0s

	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	//	r = &test.Rr{"this is a test"}
	//	fmt.Printf("%T %v\n", r, r)
	//	fmt.Println(download(r))
	//fmt.Println(download(r))
}
func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	//Type Switch 判断类型
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Content:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)

	}
}
