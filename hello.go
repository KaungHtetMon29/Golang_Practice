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
	var arr [5]int
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i, k := range arr {
		fmt.Println("index", i)
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

func sliceFun() {
	var s []string
	fmt.Println(s)

	s = make([]string, 3)
	fmt.Println(s)

	s[0] = "0"
	s[1] = "1"
	s[2] = "2"
	fmt.Println("set:", s)
	s = append(s, "3")
	fmt.Println("appended", s)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	//slice operator
	sliced := s[0:1]
	fmt.Println("slice1", sliced)
	sliced1 := s[2:]
	fmt.Println("sl2:", sliced1)
	sliced2 := s[:3]
	fmt.Println("sl3,", sliced2)
}

func mapfunc() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 2
	_, prs := m["k2"]
	fmt.Println("map:", m)
	fmt.Println("check that doeesn't exist", m["k3"])
	delete(m, "k2")
	clear(m)
	fmt.Println(prs)
}

func functest1(a int, b int) int {
	return a + b
}
func functest2(a, b int) int {
	return a + b
}
func multireturn() (int, int) {
	return 1, 2
}

func variadicfunc(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main() {
	fmt.Println("Hello world")
	fmt.Println(os.Getenv("USER"), ", Let's be friends!")
	fmt.Println(s)
	testVariables()
	fmt.Println("loop")
	loop()
	fmt.Println("condition")
	condition()
	switchCase()
	array()
	fmt.Println("Slice")
	sliceFun()
	fmt.Println("Map")
	mapfunc()
	fmt.Println("Func")
	fmt.Println(functest1(1, 2))
	fmt.Println(functest2(2, 4))
	a, b := multireturn()
	fmt.Println(a, b)
	fmt.Println("Variadic Func")
	variadicfunc(1, 2, 3, 4)
}
