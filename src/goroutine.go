package main

import (
	"time"
	"fmt"
)

func worker(n int) {
	i := 1
	for {
		fmt.Printf("%d : %d\n", n, i)
		time.Sleep(time.Duration(n) * time.Second)
		i++
	}
}

func main() {
	go worker(3)
	go worker(5)
	time.Sleep(time.Duration(60) * time.Second)
}
