package main

import (
	"godemo1/errhanding/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

// 处理错误
func errWrapper(handler appHandler) func(writer http.ResponseWriter,
	request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			//log.Warn不是标准库的
			log.Printf("Error handling request: %s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err): //不存在
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden //无权限
			default:
				code = http.StatusInternalServerError
			}
			//第一个参数汇报给谁err，第二个参数错误信息，第三个错误码
			//第二个参数如果使用err.Error()会暴露内部错误信息
			http.Error(writer,
				http.StatusText(code),
				code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
