package main

import (
	"bufio"
	"fmt"
	"godemo1/functional/fib"
	"io"
	"strings"
)

// 函数类型变量 只要是一个类型就可以实现接口
type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	//下一个数
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}

	s := fmt.Sprintf("%d\n", next)

	//p可能太小装不下一个数 解决办法把用struct 把strings.NewReader和intGen缓存起来TODO
	return strings.NewReader(s).Read(p)

}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	//省略其他条件只剩结束条件类似while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = fib.Fibonacci()
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())

	printFileContents(f)

}
