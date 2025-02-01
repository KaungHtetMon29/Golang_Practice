package main

import "fmt"

type animal interface {
	shout()
}

type human struct {
	name string
	age  int
}

func (h human) shout() {
	fmt.Println("Ahhhhh")
}

func action(animal animal) {
	animal.shout()
}

func main() {
	fmt.Println("Interface")
	h := human{"kaung", 12}
	action(h)
}
