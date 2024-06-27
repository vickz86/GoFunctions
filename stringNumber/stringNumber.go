package stringnumber

import (
	"fmt"
	"strconv"
	"unicode"

	sliceint "github.com/vickz86/GoFunctions/sliceInt"
)

// SliceIndexConnectedNumbers : return first and last index of connected letter => ab101ab54 =[2,4]
func SliceIndexConnectedNumbers(theString string) []int {
	//variable

	//  slice int with connected first number
	var sliceInt2 []int

	// get the index of all the number in the string
	var sliceInt []int

	//return slice Int
	var returnSliceInt = []int{}

	//loop through all rune of the string , get all index of numbers
	for index, letter := range theString {
		if unicode.IsNumber(letter) {
			sliceInt = append(sliceInt, index)
		}

	}

	//if sliceInt == 0 : no numbers print error and return 0
	if len(sliceInt) == 0 {
		fmt.Println("No number in inputed string")
		return []int{}
	}

	//loop through sliceInt slice and add only numbers if there are next in numbers
	for index, number := range sliceInt {
		// add first number
		if index == 0 {
			sliceInt2 = append(sliceInt2, number)
			continue
		} else if /* check if next index is separated by one */ (number - 1) == sliceInt[index-1] {
			sliceInt2 = append(sliceInt2, number)
		} else {
			break
		}

	}

	//if sliceInt2 > 2 , only keep first and last
	returnSliceInt = sliceint.KeepFirstLastSlice(sliceInt2)

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

// StringToInt : convert a string to an int
// add toAdd to the newly created int
func StringToInt(stringInt string, toAdd int) int {

	//convert input string to newInt
	newInt, err := strconv.Atoi(stringInt)
	//check if error
	if err != nil {
		fmt.Println("Error with input String!")
	}

	//add to newInt toAdd int
	newInt += toAdd

	//return
	return newInt

}

//------COMBINE PREVIOUS FUNCTION-------//

// AddToIntStringInString : add an amount to first grourp
// of connected number in a string
// example : all17boa  -> all18boa
func AddToIntStringInString(inString string, toAdd int) string {

	// first get slice of first int in string
	sliceInt := SliceIndexConnectedNumbers(inString)
	//if len sliceInt is 0 , print it and return 0
	if len(sliceInt) == 0 {
		fmt.Println("not number is string!")
		return ""
	}

	//create a before and after string using GetStringBeforeAfter
	beforeStr := GetStringBeforeAfterIndex(inString, sliceInt[0], 0)
	afterStr := GetStringBeforeAfterIndex(inString, sliceInt[len(sliceInt)-1], 1)

	//create a string for the int
	stringInt := GetStringBetweenslice(inString, sliceInt)

	//convert to int and add , using StringToInt function
	newInt := StringToInt(stringInt, toAdd)

	//convert to string newInt
	intBackString := strconv.Itoa(newInt)

	//create the final string
	finalString := beforeStr + intBackString + afterStr

	//return finalstring
	return finalString

}
