package main

import (
	"fmt"
	"os"
)

const s string = "constant"

func testVariables() {
	var a = "test"
	var b, c int = 1, 2
	var d = true
	var e int
	f := "apple"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

func loop() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for k := range 3 {
		fmt.Println("range", k)
	}
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("n modulo", n)
	}
}

func main() {
	fmt.Println("Hello world")
	fmt.Println(os.Getenv("USER"), ", Let's be friends!")
	fmt.Println(s)
	testVariables()
	loop()
}
