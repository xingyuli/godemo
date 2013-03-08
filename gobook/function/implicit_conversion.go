package main

import "fmt"

type Number int

func (n *Number) inc() {
	fmt.Println("n in inc:", n)
	*n++
}

func (n Number) print() {
	fmt.Println("&n in print:", &n)
	fmt.Println("The number is equal to", n)
}

func main() {
	i := Number(10)
	fmt.Println("i is equal to", i)

	fmt.Println("Let's increment it twice")
	i.inc()
	fmt.Println("i is equal to", i)
	(&i).inc()
	fmt.Println("i is equal to", i)

	p := &i
	fmt.Println("Let's print it twice")
	p.print()
	i.print()
}
