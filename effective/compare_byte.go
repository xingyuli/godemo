package main

import "fmt"

// Compare returns an integer comparing the two byte slices,
// lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	} 

	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}

	return 0
}

func main() {
	a := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	c := []byte{1, 'c', 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Println("a == b:", Compare(a, b))
	fmt.Println("a == c:", Compare(a, c))
}