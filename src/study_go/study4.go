package main

import (
	"fmt"
)

type Human4 struct {
	name  string
	age   int
	phone string
}

type Student4 struct {
	Human4
	school string
}

type Employee4 struct {
	Human4
	company string
}

func (h *Human4) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s \n", h.name, h.phone)
}

func main() {
	mark := Student4{Human4{"Mark", 25, "222-222-yyyy"}, "MIT"}
	sam := Employee4{Human4{"Sam", 45, "111-8888-xxxx"}, "Golang inc"}

	mark.SayHi()
	sam.SayHi()
}
