package main

import "fmt"

func isOdd(integer int) bool {
	return integer%2 != 0
}

func isBiggerThan4(integer int) bool {
	return integer > 4
}

/*
filter_factory
input: a criteria function of type: func(int) bool
output: a "spliting" function of type: func([]int) (yes, no []int)
--
To wrap your head around the declaration below, consider it this way:
// declare parameter and result types
type test_int func(int) bool
type slice_split func([]int) ([]int, []int)
and then declare it like this:
func filter_factory(f test_int) slice_split
--
In summary: filter_factory takes a function, and creates another one
of a completly differenty type.
*/
func filter_factory(f func(int) bool) func(s []int) (yes, no []int) {
	return func(s []int) (yes, no []int) {
		for _, value := range s {
			if f(value) {
				yes = append(yes, value)
			} else {
				no = append(no, value)
			}
		}
		return // we don't have to add yes, no. They're named results.
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("s = ", s)

	odd_even_function := filter_factory(isOdd)
	odd, even := odd_even_function(s)
	fmt.Println("odd =", odd)
	fmt.Println("even =", even)

	bigger, smaller := filter_factory(isBiggerThan4)(s)
	fmt.Println("Bigger than 4:", bigger)
	fmt.Println("Smaller than or equal to 4:", smaller)
}
