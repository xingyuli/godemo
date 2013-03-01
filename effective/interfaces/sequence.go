package main

import (
	"fmt"
	"sort"
)

// define a named type
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
	sort.Sort(s)
	return fmt.Sprint([]int(s))
}

func main() {
	originalSequence := Sequence([]int{13, 5, 232, 53, 21, 62, 66})
	fmt.Println([]int(originalSequence))

	asString := fmt.Sprint(originalSequence)
	fmt.Println([]int(originalSequence))
	
	fmt.Println("asString:", asString)
}
