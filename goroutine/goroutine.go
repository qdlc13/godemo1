package main

import (
	"fmt"
	"runtime"
	"time"
)

// main函数本身就是一个goroutine
func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		//内部之间并发执行 和外部main函数也是并发关系 main推出则所有协程推出
		//必须传参数否则会越界 不传参当for结束后time.Sleep(time.Millisecond)时i=10 由于是协程a[i]++会
		//引用a[10]++ 越界
		go func(i int) { //race condition
			for {
				a[i]++            //没有协程切换 不交出控制权
				runtime.Gosched() //手动交出控制权
				//fmt.Printf("hello from goroutine %d\n", i) io操作有协程之间切换
			}
		}(i)
	}
	//睡眠使得main慢点退出
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
