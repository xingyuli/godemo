package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := x
	for i := 0; i < 10; i++ {
		z = z - (cmplx.Pow(z, 3) - x) / (3 * cmplx.Pow(z, 2))
	}
	return z
}

func main() {
	fmt.Println("my:", Cbrt(2))
	fmt.Println("cmplx.Pow:", cmplx.Pow(complex128(2), 1.0 / 3))
}
