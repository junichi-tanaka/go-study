package main

import (
	"fmt"
)

func main() {
	for i := 1; true; i++ {
		if (i % (3 * 5) == 0) {
			fmt.Println("Fizz Buzz")
		} else if (i % 3 == 0) {
			fmt.Println("Fizz")
		} else if (i % 5 == 0) {
			fmt.Println("Buzz")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
