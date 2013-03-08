package main

// import "fmt"
// import "os"

import (
	"fmt"
	"os"
)

// var _ = fmt.Println
// var _ = os.Stderr

var (
	_ = fmt.Println
	_ = os.Stderr
)

const (
	w = iota
	x
	y
	z
)

const (
	a = 10 * iota
	b
	c
	d
	e
	f
	g
)

func main() {
	fmt.Println(w, x, y, z)
	fmt.Println(a, b, c, d, e, f, g)
}
