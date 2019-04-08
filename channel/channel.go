package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)

	go func() {
		for {
			fmt.Printf("Worker %d received %c\n",
				id, <-c)
		}
	}()

	return c
}

func chanDemo() {
	//var c chan int	// c == nil
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		//channels[i] = make(chan int)
		//go worker(i, channels[i])
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
}

func bufferChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	close(c)
}

func main() {
	chanDemo()
	//bufferChannel()
	//channelClose()
	time.Sleep(time.Millisecond)
}
