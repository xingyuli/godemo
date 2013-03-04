package main

import "io"
import "fmt"

var _ = fmt.Println
var _ io.Reader

func main() {
	greeting := "hello, world"
	_ = greeting
}