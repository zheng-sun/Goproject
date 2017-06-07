package main

import "fmt"

type Human5 struct {
	name string
	age int
	phone string
}

type student5 struct {
	Human5
	school string
}

type Employee5 struct {
	Human5
	company string
}

func (h *Human5) SayHi()  {
	fmt.Printf("Hi I am %s you can call me on %s \n", h.name, h.phone)
}

func (e *Employee5)  SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. call me on %s \n",e.name,e.company,e.phone)
}

func main() {
	mark := student5{Human5{"sunzheng",22,"iphone"},"石家庄"}
	sam := Employee5{Human5{"gaoyu",21,"xiaomi"},"北京"}

	mark.SayHi()
	sam.SayHi()
}
