package main

import (
	"fmt"
	"log"
)

var timeZone = map[string] int {
	"UTC": 0*60*60,
	"EST": -5*60*60,
	"CST": -6*60*60,
	"MST": -7*60*60,
	"PST": -8*60*60,
}

func Offset(tz string) int {
	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	log.Println("unknown time zone:", tz)
	return 0
}

func main() {
	fmt.Println(timeZone)
	fmt.Printf("%#v\n", timeZone)

	// fetch a map value with a key
	offset := timeZone["EST"]
	fmt.Println(offset)

	attended := map[string] bool {
		"Ann": true,
		"Joe": true,
	}

	// if the key is not present in the map, then the
	// zeroed value of its value type will be returned
	for _, person := range []string{"John", "Ann"} {
		if attended[person] {	// will be false if person is not in the map
			fmt.Println(person, "was at the meeting")
		}
	}

	// as the zero value of int type is 0, we cannot
	// distinguish them using following approach
	timeZone["ABC"] = 0
	tzNames := []string{"ABC", "notexist"}
	for _, tz := range tzNames {
		fmt.Printf("timeZone[\"%v\"] = %v\n", tz, timeZone[tz])
	}
	
	// instead, we need to use Multiple Assignment
	var seconds int
	var ok bool
	for _, tz := range tzNames {
		seconds, ok = timeZone[tz]
		fmt.Printf("timeZone[\"%v\"] = %v, existence = %v\n", tz, seconds, ok)
	}

	fmt.Println(Offset("ABC"), Offset("notexist"))

	// delete a map entry
	delete(timeZone, "PST")
	delete(timeZone, "notexist")	// it's ok to delete a not presented key
	fmt.Println(timeZone)
}
