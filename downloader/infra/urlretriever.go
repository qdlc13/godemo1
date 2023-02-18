package infra

import (
	"io"
	"net/http"
)

type Retriever struct {
}

func (r Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, _ := io.ReadAll(resp.Body)
	return string(bytes)
}
