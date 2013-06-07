package main

import (
	"fmt"
)

type Namer interface {
	Name() string
}

type Weak string

func (weak Weak) Name() string {
	return "[Weak] " + string(weak)
}

func (weak Weak) String() string {
	return weak.Name()
}

func main() {
	var w Namer = Weak("urh")
	switch w.(type) {
	case Weak:
		fmt.Println("Is weak:", w)
	default:
		fmt.Println("Unkown:", w)
	}
}
