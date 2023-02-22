package filelisting

import (
	"io"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

// 这个类型要想实现type userError interface接口,就需要完成接口中的方法
type userError string

func (e userError) Error() string {
	return e.Message()
}
func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error {
	if strings.Index(
		request.URL.Path, prefix) != 0 {
		//
		return userError("path must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):] //去掉list解析后面文件的路径
	file, err := os.Open(path)
	if err != nil {
		//panic(file)
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		//return
		return err
	}
	defer file.Close()
	all, err := io.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}
	writer.Write(all) //内容写入http.ResponseWriter
	return nil
}
