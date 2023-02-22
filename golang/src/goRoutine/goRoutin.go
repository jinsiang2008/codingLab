package main

import (
	"fmt"
	"time"
)

func say(s string){
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sc(ch chan<- string){
	ch <- "hello from ch2"
}

func speed1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "speed 1"
}

func speed2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "speed 2"
}

func main() {
	go say("world")
	say("hello")
	// say("The end")

	c := make(chan string)
	go func(){c <- "hello from c"}()
	msg := <- c
	fmt.Println(msg)
	time.Sleep(100 * time.Millisecond)

	//单向通道
	ch := make(chan string)
	go sc(ch)
	fmt.Println(<-ch)

	//select处理多通道
	c1 := make(chan string)
	c2 := make(chan string)
	go speed1(c1)
	go speed2(c2)
	fmt.Println("The first to arrive is:")
	select {
	case s1 := <- c1:
		fmt.Println(s1)
	case s2 := <- c2:
		fmt.Println(s2)
	}

	//缓冲通道
	ch2 := make(chan string, 2)
	ch2 <- "1"
	ch2 <- "2"
	// ch2 <- "3"

	fmt.Println(<-ch2)
}