package main

import (
	"log"
)

type Rect struct {
	a, b          float64
	width, height float64
}

func main() {
	r1 := new(Rect)
	r2 := &Rect{}
	r3 := &Rect{0, 0, 100, 200}
	r4 := &Rect{width: 100, height: 200}

	log.Println("r1:", r1)
	log.Println("r2:", r2)
	log.Println("r3:", r3)
	log.Println("r4:", r4)

	r5 := r4
	r5.a = 4
	log.Println("r5:", r5)
	log.Println("r4:", r4)
}
