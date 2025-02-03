package main

import (
	"fmt"
	"io"
	"os"
)

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
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
