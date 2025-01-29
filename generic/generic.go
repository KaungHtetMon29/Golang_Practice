package main

import "fmt"

type MyInt int

func testGeneric[T ~int](num T) {
	fmt.Println(num)
}

func main() {
	var testInt MyInt = 2
	testGeneric(1)
	testGeneric(testInt)
}
