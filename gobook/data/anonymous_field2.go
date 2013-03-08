package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // an anonymous field of type Human
	Skills     // an anonymous field for his skills
	int        // we will use this int as an anonymous field for his preferred number
	speciality string
}

func main() {
	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology"}

	fmt.Println("Her name is", jane.name)
	fmt.Println("Her age is", jane.age)
	fmt.Println("Her weight is", jane.weight)
	fmt.Println("Her speciality is", jane.speciality)

	jane.Skills = []string{"anotomy"}
	fmt.Println("Her skills are", jane.Skills)

	fmt.Println("She acquired two new ones")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are", jane.Skills)

	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
}
