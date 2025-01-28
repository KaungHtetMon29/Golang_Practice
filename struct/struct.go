package main

import "fmt"

type person struct {
	name string
	age  int
}

type actions interface {
	shout() string
	walk() string
}
type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
)

func (r *person) outputName() string {
	return r.name
}
func (r person) shout() string {
	return "do shout"
}

func (r person) walk() string {
	return "do walk"
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 22
	return &p
}

func main() {
	// var test int = 1
	var p person = person{"khm", 22}
	var action actions
	action = p
	fmt.Println(action.shout())
	fmt.Println(&p.name)
	fmt.Println(&p.age)
	fmt.Println(p.outputName())
	fmt.Println(Monday)
}
