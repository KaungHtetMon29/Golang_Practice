package main

import (
	"fmt"
	"os"
	"time"
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
func condition() {
	if 7%2 == 0 {
		fmt.Println("divisible")
	} else {
		fmt.Println("not divisible by 2")
	}
	if 8%4 == 0 {
		fmt.Println("divisible by 4")
	}

}

func switchCase() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend")
	default:
		fmt.Println("Its weekday")
	}
}
func array() {
	var arr [5]int
	fmt.Println(len(arr))
	//let compiler count the length of array
	var barr = [...]int{1, 2, 3}
	fmt.Println(barr)
	for idx := range len(barr) {
		fmt.Println(barr[idx])
	}
}
func main() {
	fmt.Println("Hello world")
	fmt.Println(os.Getenv("USER"), ", Let's be friends!")
	fmt.Println(s)
	testVariables()
	loop()
	condition()
	switchCase()
	array()
}
