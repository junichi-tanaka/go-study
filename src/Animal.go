package main

import (
	"fmt"
)

type Animal interface {
	Say() string
}

type Dog struct {
}

func (this *Dog) Say() string {
	return "Vow wow"
}

func main() {
	var animal Animal
	animal = new(Dog)
	fmt.Println(animal.Say())
}
