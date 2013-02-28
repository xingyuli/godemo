package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add1(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("add:", add(42, 13))
	fmt.Println("add1:", add1(11, 22))
}
