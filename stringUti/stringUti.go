package stringuti

import (
	"errors"
	"fmt"
	"strings"
)

// SliceStringIndex splits a string at the specified index and returns two substrings: before and after.
func SliceStringIndex(theString string, index int) (string, string, error) {
	// Check if the index is valid
	if index < 0 || index >= len(theString) {
		return "", "", errors.New("index is out of range")
	}
	
	// Create the two return strings
	beforeString := theString[:index+1]
	afterString := theString[index+1:]
	
	return beforeString, afterString, nil
}


// slice the string at the first meet specified letter (byte) , if more than time the element print error
func SliceStringString(theString,letter string) (string,string){

	// count the number of time the string is present and print message
	count := strings.Count(theString,letter)
	if count>1{
		fmt.Printf("CAREFULL , the letter %s is present %v time in %s",letter,count,theString)
	}
	
	_,_,isPresent := strings.Cut(theString,letter)
	if !  isPresent{
		fmt.Printf("CAREFULL , the letter %s is not present in %s",letter,theString)
	}



	return "",""

}

