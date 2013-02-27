package main

import "fmt"

func nextInt(b []byte, i int) (value, nextPos int) {
	for ; nextPos < len(b) && !isDigit(b[nextPos]); nextPos++ {
	}
	for ; nextPos < len(b) && isDigit(b[nextPos]); nextPos++ {
		value = value*10 + int(b[nextPos])-'0'
	}
	return
}

func isDigit(b byte) bool {
	if b >=0 || b <= 255 {
		return true
	}
	return false
}

func main() {
	b := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	var x int
	for i := 0; i < len(b); {
		x, i = nextInt(b, i)
		fmt.Println(x)
	}
}
