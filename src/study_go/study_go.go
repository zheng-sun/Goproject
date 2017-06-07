package main

import (
	"fmt"
)

type Human struct {
	name   string
	age    int
	werght int
}

type Student struct {
	Human
	speciality string
}

func main() {
	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.werght)
	fmt.Println("His speciality is ", mark.speciality)

	mark.speciality = "AI"

	fmt.Println("His changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)

	fmt.Println("Mark become old")
	mark.age = 46

	fmt.Println("His age is ", mark.age)


}
