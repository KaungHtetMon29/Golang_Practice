package main

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Test")
	// t.Errorf("Error")
	// _, err := os.ReadFile("test")
	// if os.IsNotExist(err) {
	// 	t.Errorf("File not found")
	// }
	// t.Log("File found")
}

func TestWrite(t *testing.T) {
	t.Log("TestWrite")
	FileWrite()
	_, err := os.ReadFile("test1")
	if os.IsNotExist(err) {
		t.Errorf("File not found")
	}
}
