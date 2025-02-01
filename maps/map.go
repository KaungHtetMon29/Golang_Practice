package main

import "fmt"

type attributes struct {
	name string
	age  int
}

func main() {
	fmt.Println("Maps")
	person := map[string]attributes{
		"kaung": {name: "KHM", age: 12},
	}
	fmt.Println(person)
	for att, value := range person {
		fmt.Println(att)
		fmt.Println(value)
	}
}
