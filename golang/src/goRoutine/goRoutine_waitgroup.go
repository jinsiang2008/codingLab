package main

import (
	"fmt"
	"sync"
)

// Goroutine 知识： https://www.liwenzhou.com/posts/Go/concurrence/
var wg sync.WaitGroup

func say(s string){
	fmt.Println(s)
	wg.Done()
}

func main() {
	wg.Add(1)
	go say("world")
	fmt.Println("hello")
	wg.Wait()
}