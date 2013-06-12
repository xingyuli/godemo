package main

import (
	"fmt"
	"time"
)

func timeout(limit time.Duration) chan bool {
	ch := make(chan bool, 1)
	go func() {
		time.Sleep(limit)
		ch <- true
	}()
	return ch
}

func main() {
	ch := make(chan int, 1)
	select {
	case <-ch:
	case <-timeout(time.Second):
		fmt.Println("timeout!!")
	}
}
