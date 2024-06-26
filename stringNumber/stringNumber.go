package stringnumber

import (
	"unicode"
)

// SliceIndexConnectedNumbers : return first and last index of connected letter => ab101ab54 =[2,4]
func SliceIndexConnectedNumbers(theString string) []int {
	//variable

	// return slice int
	var returnSliceInt []int

	// get the index of all the number in the string
	var sliceInt []int

	//loop through all rune of the string , get all index of numbers
	for index, letter := range theString {
		if unicode.IsNumber(letter) {
			sliceInt = append(sliceInt, index)
		}

	}

	//loop through sliceInt slice and add only numbers if there are next in numbers
	for index, number := range sliceInt {
		// add first number
		if index == 0 {
			returnSliceInt = append(returnSliceInt, number)
			continue
		} else if /* check if next index is separated by one */ (number - 1) == sliceInt[index-1] {
			returnSliceInt = append(returnSliceInt, number)
		} else {
			break
		}

	}

	return returnSliceInt

}

// GetStringBeforeAfterIndex , return the string before or after a specifc index : 0 =>before , 1 => after
func GetStringBeforeAfterIndex(theString string, theIndex, beforeAfter int) string {
	// variable
	var newString string

	//get the sublice before the inde
	if beforeAfter == 0 {
		newString = theString[:theIndex]

	} else /* get after */ {
		newString = theString[theIndex+1:]

	}

	//return new string
	return newString

}
