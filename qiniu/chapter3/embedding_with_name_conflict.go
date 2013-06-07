package main

import (
	"fmt"
)

type Namer struct {
	Name string
}

type Room struct {
	Namer
	Name string
}

func main() {
	room := &Room{
		Namer: Namer{"in Namer"},
		Name:  "in Room",
	}
	fmt.Println("room.Name:", room.Name)
	fmt.Println("room.Namer.Name:", room.Namer.Name)
}
