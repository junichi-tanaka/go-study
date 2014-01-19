package main

import (
	"fmt"
	"time"
)

func worker(out chan int) {
	time.Sleep(time.Second * time.Duration(10))
	out <- 10
}

func main() {
	ch := make(chan int)
	go worker(ch)
	result := <- ch
	fmt.Println(result)
}
