package main

import (
	"fmt"
	"sync"
)

var (
	a    string
	once sync.Once
)

func setup() {
	a = "hello, world"
}

func doPrint() {
	once.Do(setup)
	fmt.Println("a:", a)
}

func printTwice(quit chan bool) {
	go doPrint()
	go doPrint()
	quit <- true
}

func main() {
	quit := make(chan bool)
	go printTwice(quit)
	<-quit
}
