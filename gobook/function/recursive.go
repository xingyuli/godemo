package main

import (
	"fmt"
)

func Invert(slice []int) {
	length := len(slice)
	if length > 1 {
		slice[0], slice[length-1] = slice[length-1], slice[0]
		Invert(slice[1 : length-1])
	}
}

func Sum(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	return slice[0] + Sum(slice[1:])
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice =", slice)
	Invert(slice)
	fmt.Println("Invert(slice) =", slice)

	fmt.Println("Sum(slice) =", Sum(slice))
}
