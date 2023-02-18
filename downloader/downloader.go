package main

import (
	"fmt"
	"godemo1/downloader/infra"
)

func getRetriever() retriever {
	return infra.Retriever{}
}

type retriever interface {
	Get(string) string
}

func main() {

	var r retriever = getRetriever()
	text := r.Get("http://www.baidu.com/")
	fmt.Println(text)

}
