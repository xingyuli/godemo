package main

import "fmt"

func main() {
	for pos, char := range "日本语" {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}

	a := []byte{0, 1, 2, 4, 8, 16, 32, 64}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)

	c := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	for i, j := 0, len(c)-1; i < j; i , j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}
	fmt.Println(c)
}
