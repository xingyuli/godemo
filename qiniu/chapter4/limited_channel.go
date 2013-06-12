package main

import "fmt"

func main() {
	both := make(chan int, 10)
	read_only := (<-chan int)(both)
	write_only := (chan<- int)(both)

	for i := 0; i < 10; i++ {
		write_only <- i
	}
	close(write_only)

	for value := range read_only {
		fmt.Println("Received:", value)
	}
}
