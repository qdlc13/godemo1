package main

import "fmt"

func adder() func(int) int {
	sum := 0                 //自由变量 环境
	return func(v int) int { //v 变量
		sum += v
		return sum
	}
}

// 正统函数式编程 没有环境自由变量
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder()
	for i := 0; i <= 10; i++ {
		fmt.Printf("0+...+%d=%d\n", i, a(i))
	}

	b := adder2(0)
	for i := 0; i <= 10; i++ {
		var s int
		s, b = b(i)
		fmt.Printf("0+...+%d=%d\n", i, s)
	}

}
