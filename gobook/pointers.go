package main

import "fmt"

var (
	i     int       = 9
	hello string    = "Hello world"
	pi    float32   = 3.14
	c     complex64 = 3 + 5i
)

func Pointers() {
	fmt.Println("Hexadecimal address of i is:", &i)
	fmt.Println("Hexadecimal address of hello is:", &hello)
	fmt.Println("Hexadecimal address of pi is:", &pi)
	fmt.Println("Hexadecimal address of c is:", &c)

	var iPtr *int
	var helloPtr *string

	iPtr = &i
	helloPtr = &hello

	fmt.Println("iPtr:", iPtr)
	fmt.Println("helloPtr:", helloPtr)

	fmt.Println(*iPtr == i)
	fmt.Println(iPtr == &i)
}

func Sum() {
	sum := 0
	var double_sum *int
	for i := 0; i < 10; i++ {
		sum += i
	}
	double_sum = new(int)
	*double_sum = sum * 2
	fmt.Println("The sum of numbers from 0 to 10 is:", sum)
	fmt.Println("The double of this sum is:", *double_sum)
}

func main() {
	Pointers()
	Sum()
}
