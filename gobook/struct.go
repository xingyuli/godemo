package main

import "fmt"

type person struct {
	name string
	age  int
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var tom person
	tom.name, tom.age = "Tom", 18

	bob := person{age: 25, name: "Bob"}
	paul := person{"Paul", 43}

	tbOlder, tbDiff := Older(tom, bob)
	tpOlder, tpDiff := Older(tom, paul)
	bpOlder, bpDiff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tbOlder.name, tbDiff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, paul.name, tpOlder.name, tpDiff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		bob.name, paul.name, bpOlder.name, bpDiff)
}
