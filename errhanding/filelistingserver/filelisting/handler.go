package filelisting

import (
	"io"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error {
	path := request.URL.Path[len("/list/"):] //去掉list
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
