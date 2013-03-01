package main

import (
	"fmt"
	"os"
)

type T struct {
	a int
	b float64
	c string
}

func (t *T) String() string {
	return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}

func Min(a ...int) int {
	min := int(^uint(0) >> 1)	// largest int
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func TestFmt(text string, v interface{}) {
	fmt.Println(">>>", text)
	fmt.Printf("%q\n", v)
	fmt.Printf("%#q\n", v)
	fmt.Printf("%x\n", v)
	fmt.Printf("%T\n", v)
}

func main() {
	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))

	var x uint64 = 1<<64 - 1
	fmt.Printf("%d %x; %d, %x\n", x, x, int64(x), int64(x))

	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)

	TestFmt("string",		"abcdef")
	TestFmt("[...]byte",	[...]byte{'a', 'b', 'c', 'd', 'e', 'f'})
	TestFmt("[]byte",		[]byte{'a', 'b', 'c', 'd', 'e', 'f'})

	fmt.Println(Min(1, 4, 5, 2, -3, 8, -2, 10))
}
