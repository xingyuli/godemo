package main

import "fmt"
import "math"

func Sqrt(x float64) float64 {
	z, temp := x, 0.0
	for math.Abs(temp - z) > 1e-9 {
		temp = z
		z = z - (z * z - x) / (2 * z)
	}
	return z
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("%g, %g\n", Sqrt(float64(i)), math.Sqrt(float64(i)))	
	}
}
