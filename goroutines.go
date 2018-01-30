package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	go f("direct")

	go f("golang")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

	//channels
	message := make(chan string)

	go func() { message <- "ping" }()

	msg := <-message
	fmt.Println(msg)

	//channel buffering
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	//channel sync
	done := make(chan bool, 1)
	go worker(done)

	<-done

	//channel directions
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}
