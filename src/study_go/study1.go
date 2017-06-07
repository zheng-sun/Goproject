package main

import "fmt"

type Human1 struct {
	name string
	age string
	phone string
}

type Employyee1 struct {
	Human1
	speciality string
	phone string

}

func main() {
	Bob := Employyee1{Human1{"Bob", "34", "777-444-xxxx"},"Designer","333-222"}
	fmt.Println("Bob's work phone is:",Bob.phone)
	fmt.Println("Bob's personal phone is:",Bob.Human1.phone)
}