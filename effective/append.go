package main

import "fmt"

func debug(s []int) {
	fmt.Printf("slice: %v, len: %d, cap: %d\n", s, len(s), cap(s))
}

func Sum(v ...int) (sum int){
	fmt.Printf("%T\n", v)
	for _, e := range v {
		sum += e
	}
	return
}

func main() {
	x := []int{1, 2, 3}
	debug(x)

	x = append(x, 4, 5, 6)
	debug(x)

	// use "..." to treat a slice as a list of arguments
	y := []int{4, 5, 6}
	x = append(x, y...)
	debug(x)

	fmt.Println("sum:", Sum(x...))
}
