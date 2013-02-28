package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, ErrNegativeSqrt(f)
	}

	z, temp := f, 0.0
	for math.Abs(temp - z) > 1e-9 {
		temp = z
		z = z - (z * z - f) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
