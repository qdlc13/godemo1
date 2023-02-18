package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

// 如果struct很大值传递拷贝费时间，采用指针类型
func (r *Retriever) Get(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)

	resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)

}
