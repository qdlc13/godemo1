package main

import (
	"bufio"
	"errors"
	"fmt"
	"godemo1/functional/fib"
	"os"
)

// defer后面的语句在return和panic之前返回 而且返回顺序类似栈先进后出 先打印2然后是1
// 参数在defer语句时计算
func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	//panic("ERR occurred")
	return

	fmt.Println(4)
}
func tryDefer2() {
	for i := 0; i < 100; i++ {
		defer fmt.Printf("i=%d\n", i)
		if i == 30 {
			panic("printed too many")
		}
	}

}
func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	//close文件
	defer file.Close()
	//直接写文件慢 用bufio包装先写到内存中，到达一定数量一下子写到文件中
	//写到bufio中
	writer := bufio.NewWriter(file)
	//函数结束写入文件中
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		//把第二个参数返回值写入io.writer
		fmt.Fprintln(writer, f())

	}
}
func writeFile2(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	//if err != nil {
	//	fmt.Println("Error:", err.Error())
	//	return
	//}
	err = errors.New("this is a custom error")
	if err != nil {
		//err.(*os.PathError)判断类型
		if PathError, ok := err.(*os.PathError); !ok {
			//不是pathError
			panic(err)
		} else {
			fmt.Printf("%s    %s    %s", PathError.Op, PathError.Path,
				PathError.Error())
		}
		return
	}
	//close文件
	defer file.Close()
	//直接写文件慢 用bufio包装先写到内存中，到达一定数量一下子写到文件中
	//写到bufio中
	writer := bufio.NewWriter(file)
	//函数结束写入文件中
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		//把第二个参数返回值写入io.writer
		fmt.Fprintln(writer, f())

	}
}

func main() {
	tryDefer()
	writeFile("fib.txt")
	//tryDefer2()
	writeFile2("fib.txt")
}
