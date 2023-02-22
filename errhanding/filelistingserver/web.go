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
		defer func() {
			//服务中断恢复保护代码
			if r := recover(); r != nil {
				log.Printf("Panic:%v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			//log.Warn是第三方库方法不是标准库的
			log.Printf("Error handling request: %s", err.Error())
			//用户自建错误 user error
			if userError, ok := err.(userError); ok {
				http.Error(writer, userError.Message(), http.StatusBadRequest)
				return
			}
			//系统自带的错误system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err): //不存在
				code = http.StatusNotFound
			case os.IsPermission(err): //无权限
				code = http.StatusForbidden
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

type userError interface {
	error            //系统
	Message() string //用户
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
