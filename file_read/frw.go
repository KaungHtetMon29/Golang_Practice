package main

import (
	"fmt"
	"os"
)

var txt = "test"

func FileWrite() {
	os.WriteFile("test", []byte(txt), 0666)
	fmt.Println("file written")
}

func read() string {
	var value, err = os.ReadFile("../Golang_Practice/test1")
	if err != nil {
		return err.Error()
	}
	return string(value)
}

func main() {
	fmt.Println("file read and write")
	FileWrite()
	fmt.Println(read())
}
