package main

import "fmt"

type Box struct {
	name string
}

func (b *Box) SetName(newName string) {
	/*
		Go assumes we wanted *b, so it simplifies this for us implicitly!
		Thus the following code has the same effect as:
		(*b).name = newName
	*/
	b.name = newName
}

func main() {
	b := Box{"haha"}
	b.SetName("blah")
	fmt.Println(b)
}
