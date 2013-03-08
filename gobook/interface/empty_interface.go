package main

import "fmt"
import "strconv"

type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) String() string {
	return "❰" + h.name + " - " + strconv.Itoa(h.age) + " years -  ✆ " + h.phone + "❱"
}

func Example1() {
	var a interface{}
	var i int = 5
	s := "Hello world"

	a = i
	fmt.Println(a)

	a = s
	fmt.Println(a)
}

func Example2() {
	bob := Human{"Bob", 39, "000-777-XXX"}
	fmt.Println("This Human is :", bob)
}

func main() {
	Example1()
	Example2()
}
