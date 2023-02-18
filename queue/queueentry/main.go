package main

import (
	"fmt"
	"godemo1/queue"
)

func main() {

	q := queue.Queue{1}
	q.Push(4)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	q.Push("山东")
	fmt.Println(q.Pop())

	//	q.Push3("ss")
}
