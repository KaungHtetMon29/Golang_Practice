package main

import (
	"fmt"
	"os"
)

var txt = "test"

func write() {
	os.WriteFile("test", []byte(txt), 0666)
}

func read() string {
	var value, err = os.ReadFile("test")
	if err != nil {
		return err.Error()
	}
	return string(value)
}

func main() {
	fmt.Println("file read and write")
	write()
	fmt.Println(read())
}
