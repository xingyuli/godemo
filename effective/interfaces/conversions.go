package main

import (
	"fmt"
	"sort"
)

type Sequence []int

func (s Sequence) String() string {
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

func showTypeAndAddress(a interface{}) {
	fmt.Printf("%T, %v\n", a, &a)
}

func main() {
	l := Sequence([]int{1, 2, 3})
	showTypeAndAddress(l)
	showTypeAndAddress([]int(l))

	i := 3
	showTypeAndAddress(i)
	showTypeAndAddress(float32(i))

	fmt.Println(Sequence([]int{3, 2, 1, -12, 52, 112}))
}
