package stringuti

import (
	"errors"
)

// split a string at index , and return 2 string: before and after
func SliceStringIndex(theString string, index int) (string, string, error) {

	// check the index is not bigger than the len of the string
	if index > (len(theString) - 1) {
		err := errors.New("index is out of range! too high")
		return "", "", err
	}

	// create the 2 return strings
	var beforeString string
	var afterString string

	cutIndex := index + 1

	// create the 2 string
	beforeString = theString[:cutIndex]
	afterString = theString[cutIndex:]

	return beforeString, afterString, nil

}
