package main

import "fmt"

func LIFO() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}
}

func trace(s string) {
	fmt.Println("entering:", s)
}

func untrace(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	trace("a")
	defer untrace("a")

	fmt.Println("running in a...")
}

func main() {
	LIFO()
	a()
}
