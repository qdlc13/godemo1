package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) { //chan也是一等公民
	for {
		//n := <-c //chan发数据
		n, ok := <-c
		if !ok {
			break
		}

		fmt.Printf("Worker %d received %c \n", id, n)
	}
	//等价写法 //c是默认值就不会打印
	//for n := range c {
	//	fmt.Printf("Worker %d received %c \n", id, n)
	//}
}
func createWorker(id int) chan<- int { //外面只能给chan发数据 //<-chan int 外面只能从chan中获得数据
	c := make(chan int)
	//必须需要go func() 否则因为c没有数据也会一直循环挂掉
	go worker(id, c)
	return c
}

func chanDemo() {
	//var c chan int //c == nil
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		//channels[i] = make(chan int)
		channels[i] = createWorker(i)

		//go worker(i, channels[i])

	}
	//c := make(chan int)
	//go worker(0, c)
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
		//f := <-channels[i] 不可以
	} //chan收数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	} //chan收数据

	time.Sleep(time.Millisecond)
}
func bufferedChannel() {
	c := make(chan int, 3) //缓冲区是3 若没有缓冲区 给chan int数据而不取会报错
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}
func channelClose() {
	c := make(chan int, 3) //缓冲区是3 若没有缓冲区 给chan int数据而不取会报错
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c) //通知接收方发完数据了 接收方会此时收到类型的默认值
	time.Sleep(time.Millisecond)

}

func main() {

	//chanDemo()
	//bufferedChannel()
	channelClose()
}
