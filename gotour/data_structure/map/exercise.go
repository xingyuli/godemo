package main

import (
	"strings"
	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counter := make(map[string]int)
	for _, word := range words {
		counter[word] = counter[word] + 1
	}
	return counter
}

func main() {
	wc.Test(WordCount)
}