package main

import (
	"fmt"
	"time"
)

func doWorker(id int, c chan int, done chan bool) { //chan也是一等公民
	for n := range c {
		fmt.Printf("Worker %d received %c \n", id, n)

		go func() { done <- true }()
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

func chanDemo() {
	//var c chan int //c == nil
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	for i, worker := range workers {
		worker.in <- 'a' + i

	}
	for i, worker := range workers {
		worker.in <- 'A' + i

	}
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
