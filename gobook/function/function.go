package main

import "fmt"

func add1(a int) int {
	a += 1
	return a
}

func add2(a *int) {
	*a++
}

func main() {
	x := 3

	fmt.Println("x =", x)

	x1 := add1(x)
	fmt.Println("x+1 =", x1)
	fmt.Println("x =", x)

	add2(&x)
	fmt.Println("x+1 = ", x)
	fmt.Println("x = ", x)
}
