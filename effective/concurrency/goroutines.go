package main

import "fmt"
import "time"


func track(title string) string {
	fmt.Println("-> start:", title)
	return title
}

func un(title string) {
	fmt.Println("-> end:", title)
}

func Announce(message string, delay time.Duration) {
	defer un(track("Announce"))
	c := make(chan int)
	go func() {
		defer un(track("goroutine"))
		time.Sleep(delay)
		fmt.Println(message)
		c <- 1
	}()	// Note the parentheses - must call the function
	<-c
}

func main() {
	defer un(track("main"))
	Announce("Hello Xingyu!", time.Duration(3)*time.Second)
}
