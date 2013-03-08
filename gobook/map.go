package main

import "fmt"

var rating = map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}

func PrintKey(key string) {
	value, ok := rating[key]
	fmt.Printf("key: %s, value: %v, presence: %v\n", key, value, ok)
}

func main() {
	fmt.Println(rating)

	PrintKey("c#")

	for key, value := range rating {
		fmt.Printf("%s language is rated at %g\n", key, value)
	}

	fmt.Print("We rated these languages: ")
	count := 0
	for key := range rating {
		count++
		if count == len(rating) {
			fmt.Print(key, ".\n")
		} else {
			fmt.Print(key, ",")
		}
	}
}
