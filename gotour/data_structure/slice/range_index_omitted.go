package main

import "fmt"

func main() {
	pow := make([]int, 10)

	// value omitted
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	// index omitted
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
