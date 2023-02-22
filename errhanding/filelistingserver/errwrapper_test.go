package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// 表格驱动测试+函数式编程
type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}
func (e testingUserError) Message() string {
	return string(e)
}

func errPanic(writer http.ResponseWriter,
	request *http.Request) error {
	panic(123)
}

func errUserError(writer http.ResponseWriter,
	request *http.Request) error {
	return testingUserError("user error")
}
func errNotFound(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrNotExist
}
func errNoPermission(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrPermission
}
func errUnknown(writer http.ResponseWriter,
	request *http.Request) error {
	return errors.New("unknow error")
}
func noError(writer http.ResponseWriter,
	request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

// 公用测试数据
var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

// 假参数测试一段代码速度快 类似单元测试
func TestErrWrapper(t *testing.T) {

	for _, tt := range tests {
		//包装 这是需要测的目标函数的行为
		f := errWrapper(tt.h)
		//假的response,request参数
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet, "http://www.baidu.com", nil)
		f(response, request)
		//*ResponseRecorder 转*Response使用response.Result()
		verifyResponse(response.Result(), tt.code, tt.message, t)

		//b, _ := io.ReadAll(response.Body)
		////去掉b中的换行符
		//body := strings.Trim(string(b), "\n")
		//if response.Code != tt.code || body != tt.message {
		//	t.Errorf("expect(%d  %s); got(%d  %s);", tt.code, tt.message,
		//		response.Code, body)
		//}
	}
}

// 使用真server发送http请求测试 测试整个服务器 慢但是覆盖代码广
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		//newserver参数是一个http.Handler接口 使用http.HandlerFunc类型转换f函数
		//因为HandlerFunc这个类型实现了http.Handler这个接口 //新建server
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		verifyResponse(resp, tt.code, tt.message, t)
		//b, _ := io.ReadAll(resp.Body)
		////去掉b中的换行符
		//body := strings.Trim(string(b), "\n")
		//if resp.StatusCode != tt.code || body != tt.message {
		//	t.Errorf("expect(%d  %s); got(%d  %s);", tt.code, tt.message,
		//		resp.StatusCode, body)
		//}

	}
}

// 提出两种测试方法公共代码
func verifyResponse(response *http.Response, expectedCode int, expectedMsg string,
	t *testing.T) {

	b, _ := io.ReadAll(response.Body)
	//去掉b中的换行符
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("expect(%d  %s); got(%d  %s);", expectedCode, expectedMsg,
			response.StatusCode, body)
	}

}
