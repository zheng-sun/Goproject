package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s \n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s \n", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main()  {
	mike := Student{Human{"Mike", 25,"222-222-xxx"},"MIT",0.00}
	paul := Student{Human{"Paul",26,"111-222-xxx"},"Harvard",100}
	sam := Employee{Human{"Sam",36,"444-222-xxx"},"Golang inc", 1000}
	tom := Employee{Human{"Tom", 37,"222-4444-xxxx"},"Things inc",5000}

	var i Men
	i = mike
	fmt.Printf("This is Mike, a Student: \n")
	i.SayHi()
	i.Sing("November rain")

	//i能存储Student
	i = mike
	fmt.Println("This is Mike, a Student: \n")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employee
	i = tom
	fmt.Println("This is tom, an Employee: \n")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x{
		value.SayHi()
	}
}
