package main

import "fmt"

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

func main() {
	fmt.Println(string('!'), shouldEscape('!'))
	fmt.Println(string('@'), shouldEscape('@'))
	fmt.Println(string('#'), shouldEscape('#'))
	fmt.Println(string('$'), shouldEscape('$'))
	fmt.Println(string('%'), shouldEscape('%'))
	fmt.Println(string('^'), shouldEscape('^'))
	fmt.Println(string('&'), shouldEscape('&'))
	fmt.Println(string('*'), shouldEscape('*'))
	fmt.Println(string('('), shouldEscape('('))
	fmt.Println(string(')'), shouldEscape(')'))
	fmt.Println(string('-'), shouldEscape('-'))
	fmt.Println(string('+'), shouldEscape('+'))
}
