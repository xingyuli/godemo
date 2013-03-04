package main

import "fmt"
import "time"
import "sort"

type Sequence []int

func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sequence) String() string {
	return fmt.Sprint([]int(s))
}

// dummy
func doSomethingForAWhile() {
	time.Sleep(time.Duration(3)*time.Second)
}

func main() {
	c := make(chan int)	// Allocate a channel.
	list := Sequence([]int{1, 45, -1, 3, -56, 42, 14})

	fmt.Println("before:", list)
	defer fmt.Println("after:", list)

	// Start the sort in a goroutine; when it completes, signal on the channel.
	go func() {
		sort.Sort(list)
		c <- 1	// Send a signle; value does not matter.
	}()

	doSomethingForAWhile()
	<-c	// Wait for sort to finish; discard sent value.
}
