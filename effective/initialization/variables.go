package main

import (
	"fmt"
	"os"
)

var (
	home	= os.Getenv("HOME")
	user	= os.Getenv("USER")
	goRoot	= os.Getenv("GOROOT")
)

func main() {
	fmt.Printf("home: %s, user: %s, goRoot: %s", home, user, goRoot)
}
