package main

import "fmt"

type person struct {
	name string
	age  int
}

func Older(people ...person) (bool, person) {
	if 0 == len(people) {
		return false, person{}
	}

	older := people[0]
	for _, value := range people {
		if value.age > older.age {
			older = value
		}
	}
	return true, older
}

func main() {
	var (
		ok    bool
		older person
	)

	paul := person{"Paul", 23}
	jim := person{"Jim", 24}
	sam := person{"Sam", 84}
	rob := person{"Rob", 54}
	karl := person{"Karl", 19}

	_, older = Older(paul, jim)
	fmt.Println("The older of Paul and Jim is:", older.name)

	_, older = Older(paul, jim, sam)
	fmt.Println("The older of Paul, Jim and Sam is:", older.name)

	_, older = Older(paul, jim, sam, rob)
	fmt.Println("The older of Paul, Jim, Sam and Rob is:", older.name)

	_, older = Older(karl)
	fmt.Println("When Karl is alone in a group, the older is:", older.name)

	ok, older = Older()
	if !ok {
		fmt.Println("In an empty group there is no older person")
	}

	a_slice := []int{1, 2, 3}
	fmt.Println(append(a_slice, 4, 5))
	fmt.Println(a_slice)
}
