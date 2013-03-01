package main

import "fmt"

func main() {
	switch t := interfaceValue.(type) {
	default:
		fmt.Println("unexpected type %T", t)	// %T prints type
	case bool:
		fmt.Println("boolean %t\n", t)
	case int:
		fmt.Println("integer %d\n", t)
	case *bool:
		fmt.Println("pointer to boolean %t\n", *t)
	case *int:
		fmt.Println("pointer to integer %d\n", *t)
	}
}
