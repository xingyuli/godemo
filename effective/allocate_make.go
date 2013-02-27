package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("type: %T, len: %d, cap: %d\n",
		s, len(s), cap(s))
}

func printSlicePtr(sPtr *[]int) {
	fmt.Printf("type: %T, len: %d, cap: %d\n",
		sPtr, len(*sPtr), cap(*sPtr))
}

func main() {
	/*
	 * new slice
	 */
	var p *[]int = new([]int)	// allocates slice structure; *p == nil; rarely useful
	printSlicePtr(p)
	fmt.Println("*p == nil:", *p == nil)

	*p = make([]int, 100, 100)
	printSlicePtr(p)

	/*
	 * make slice
	 */
	var v []int = make([]int, 100)
	printSlice(v)
	printSlicePtr(&v)

	v2 := make([]int, 100)
	printSlice(v2)

	/*
	 * make map
	 */
	m := make(map[string]int)
	fmt.Printf("type: %T, len: %d, m == nil: %v\n", m, len(m), m == nil)
}
