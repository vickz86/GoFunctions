package slicefloat

import (
	"errors"

	singlefloat "github.com/vickz86/GoFunctions/singleFloat"
)

// CompareLastToBefore : compare last value of slice with another value
// return true if last value is bigger , float of difference
// 0 to compare with first , numbers to walk backward
func CompareLastToBefore(ToCompare []float64, CompareWith int) (bool, float64, error) {
	// check for error
	if CompareWith > len(ToCompare) {
		return false, 0, errors.New("error,CompareWith int is too big")
	}

	//variable
	// last in slice
	lastInSlice := ToCompare[len(ToCompare)-1]

	//choose with what to compare with
	var compInSLice float64

	if CompareWith == 0 {
		compInSLice = ToCompare[0]
	} else {
		compInSLice = ToCompare[len(ToCompare)-1-CompareWith]
	}

	// get the difference between the 2 value
	diff := singlefloat.Compare2Float(lastInSlice, compInSLice)

	//create bool
	var bigger bool
	if diff > 0 {
		bigger = true
	}

	return bigger, diff, nil

}
