package main

import (
	"fmt"
	"time"
)

// pinger can only send to channel
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// ponger can send and receive to channel
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// printer can only receive from channel
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
