package main

import "fmt"
import "os"

var user = os.Getenv("USER")

func init() {
	if user == "" {
		panic("no values for $USER")
	}
}

func main() {
	fmt.Println("main")
}
