package main

import (
	"fmt"
	"godemo1/retriever/mock"
	"godemo1/retriever/real"
	"time"
)

const url = "http://www.baidu.com"

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

// 接口的组合

type Poster interface {
	Post(url string, form map[string]string) string
}

// 组合
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {

	s.Post(url, map[string]string{
		"contents": "another faked baidu.com",
	})

	return s.Get(url)
}

func main() {
	//r不是简单的引用 有类型%T和值%v
	var r Retriever
	retriever := mock.Retriever{"This is a fake baidu.com"}
	r = &retriever
	//r = &mock.Retriever{"This is a fake baidu.com"}也可以
	//r = mock.Retriever{"This is a fake baidu.com"}
	fmt.Println(download(r))
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	//判断类型的等价写法
	//Type Assertion
	realRetriever := r.(*real.Retriever) //判断是不是realRetriever 是的话把值赋给前面
	fmt.Println(realRetriever.TimeOut)   //1m0s

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	//r = &test.Rr{"this is a test"}
	//fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))
	//fmt.Println(download(r))
	//fmt.Println(download(&test.Rr{Contents: "dd"}))
	//r = &mock.Retriever{"This is a fake baidu.com"}
	fmt.Println("Try a session")
	fmt.Println(session(&retriever)) //不能传r因为r只有Retriever的能力缺少方法的实现
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch:")
	//Type Switch 判断类型
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Content:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
