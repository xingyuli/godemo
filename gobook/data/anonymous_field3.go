package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string // his own mobile number
}

type Employee struct {
	Human
	speciality string
	phone      string // work's phone number
}

func main() {
	bob := Employee{Human{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}
	fmt.Println("Bob's work phone is:,", bob.phone)
	// Now we need to specify Human to access Human's phone
	fmt.Println("Bob's personal phone is:", bob.Human.phone)
}
