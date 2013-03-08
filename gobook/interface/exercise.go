package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	age  int
}

type IntSlice []int
type Float32Slice []float32
type PersonSlice []Person

type MaxInterface interface {
	Len() int
	Get(i int) interface{}
	Bigger(i, j int) bool
}

// Len implementation for our three types
func (x IntSlice) Len() int     { return len(x) }
func (x Float32Slice) Len() int { return len(x) }
func (x PersonSlice) Len() int  { return len(x) }

// Get implementation for our three types
func (x IntSlice) Get(i int) interface{}     { return x[i] }
func (x Float32Slice) Get(i int) interface{} { return x[i] }
func (x PersonSlice) Get(i int) interface{}  { return x[i] }

// Bigger implementation for our three types
func (x IntSlice) Bigger(i, j int) bool     { return x[i] > x[j] }
func (x Float32Slice) Bigger(i, j int) bool { return x[i] > x[j] }
func (x PersonSlice) Bigger(i, j int) bool  { return x[i].age > x[j].age }

func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func Max(data MaxInterface) (ok bool, max interface{}) {
	if data.Len() == 0 {
		return false, nil
	}

	max = data.Get(0)
	m := 0
	for i := 0; i < data.Len(); i++ {
		if data.Bigger(i, m) {
			max = data.Get(i)
			m = i
		}
	}
	return true, max
}

func main() {
	iSlice := IntSlice{1, 2, 44, 6, 44, 222}
	fSlice := Float32Slice{1.99, 3.14, 24.8}
	pSlice := PersonSlice{
		Person{name: "Bart", age: 24},
		Person{name: "Bob", age: 23},
		Person{name: "Gertrude", age: 104},
		Person{name: "Paul", age: 44},
		Person{name: "Sam", age: 34},
		Person{name: "Jack", age: 54},
		Person{name: "Martha", age: 74},
		Person{name: "Leo", age: 4},
	}

	_, m := Max(iSlice)
	fmt.Println("The biggest interger in iSlice is:", m)

	_, m = Max(fSlice)
	fmt.Println("The biggest float in fSlice is:", m)

	_, m = Max(pSlice)
	fmt.Println("The oldest person in pSlice is:", m)
}
