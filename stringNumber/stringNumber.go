package stringnumber

import (
	"fmt"
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

// GetStringBetweenslice : get string between element of slice
// if len slice is 0  => error
// if len slice is 1  => return single character
// if len slice == 2  => "ab12ab" : "12"
// if bigger than 2 , error
func GetStringBetweenslice(theString string, sliceIndex []int) string {
	//variable
	var returnString string

	//if slice is 0 or bigger than 2 => error
	if len(sliceIndex) == 0 || len(sliceIndex) > 2 {
		fmt.Println("error! not correct number of element is slice")
		return ""
	}

	//if slice is 1 return 1 string
	if len(sliceIndex) == 1 {
		returnString = string(theString[sliceIndex[0]])

	} else /* slice based on 2 element in slice */ {
		returnString = theString[sliceIndex[0]:(sliceIndex[1] + 1)]

	}

	return returnString

}
