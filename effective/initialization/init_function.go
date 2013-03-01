package main

import (
	"fmt"
	"log"
	"flag"
)

type MyInt int

func (i MyInt) String() string {
	fmt.Println("MyInt.String()")
	return fmt.Sprintf("this is MyInt(%d)", int(i))
}

var (
	i = fmt.Sprint(MyInt(1))
	user, home, goRoot = "liuuyxin", "", ""
)

func init() {
	fmt.Println("init...")

	if user == "" {
		log.Fatal("$USER not set")
	}
	if home == "" {
		home = "/home/" + user
	}
	if goRoot == "" {
		goRoot = home + "/go"
	}

	// goRoot may be overriden by --goroot flag on command line.
	flag.StringVar(&goRoot, "goroot", goRoot, "Go root directory")
}

func main() {
	fmt.Println("main:", i)
	fmt.Println("goRoot:", goRoot)
}
