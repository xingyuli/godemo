package main

import "fmt"

type Base struct {
	Name string
}

func (b *Base) Foo() {
	fmt.Println("base.Foo")
}
func (b *Base) Bar() {
	fmt.Println("base.Bar")
}

type Foo struct {
	Base
}

func main() {
	f := new(Foo)
	f.Foo()
	f.Bar()
}
