package main

import "fmt"

// Sum use pointer to array to avoid copying values
func Sum(a *[3]float64) (sum float64){
	for _, v := range *a {
		sum += v
	}
	return
}

func Append(slice, data []byte) []byte {
	l := len(slice)
	if l + len(data) > cap(slice) {	// reallocate
		// Allocate double what's needed, for future growth
		newSlice := make([]byte, (l + len(data))*2)
		// The copy function is predeclared and works for any slice type
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}
	return slice
}

func main() {
	/*
	 * arrays are values
	 */
	a := [...]int{1, 2, 3, 4}
	b := a
	b[0] = -1	// b: {-1, 2, 3, 4}, but a: {1, 2, 3, 4} still
	fmt.Printf("a = %v, b = %v\n", a, b)

	/*
	 * use array in function is expensive, if you really
	 * want to use an array, pass a pointer to the array
	 * to the function
	 */
	array := [...]float64{7.0, 8.5, 9.1}
	fmt.Println(Sum(&array))	// Note the explicit address-of operator

	/*
	 * slices are reference types
	 */
	s := []int{1, 2, 3, 4}
	t := s 		// t and s refer to the same underlying array
	t[0] = -1	// both s and t will be {-1, 2, 3, 4}
	fmt.Printf("s = %v, t = %v\n", s, t)

	original := []byte{'a', 'b', 'c', 'd'}
	toAppend := []byte{'e', 'f', 'g', 'h'}
	// use the self defined Append function
	fmt.Println(Append(original, toAppend))
	// use the built-in append function
	fmt.Println(append(original, 'z'))
}
