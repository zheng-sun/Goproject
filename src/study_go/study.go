package main

import "fmt"

type Skills []string

type Human0 struct {
	name string
	age int
	weight int
}
type Student0 struct {
	Human0
	Skills
	int
	speciality string
}

func main() {
	jane := Student0{Human0:Human0{"jane", 35, 100}, speciality:"Biology"}
	fmt.Println("Her name is ",jane.Human0.name)
	fmt.Println("her age is ",jane.age)
	fmt.Println("her weight is ", jane.weight)
	fmt.Println("her speciality is ", jane.speciality)

	jane.Skills = []string{"anatomy"}
	fmt.Println("her skills are ", jane.Skills)
	fmt.Println("she acquired two new ones")

	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)

	jane.int = 3
	fmt.Println("Her preferred number is ", jane.int)
}