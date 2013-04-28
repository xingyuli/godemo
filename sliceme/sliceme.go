package sliceme

import (
	"fmt"
)

type Iterator func(item interface{}, index int) interface{}

func Map(s []interface{}, iter Iterator) []interface{} {
	if s == nil {
		return nil
	}
	if iter == nil {
		return s
	}

	var retVal []interface{}
	for index, item := range s {
		retVal = append(retVal, iter(item, index))
	}
	return retVal
}
