package stringnumber

import (
	"fmt"
	"unicode"
)

// SliceIndexConnectedNumbers : return first and last index of connected letter => ab101ab :[2,4]
func SliceIndexConnectedNumbers(theString string) /*  []int */ {
	//loop through all rune of the string
	for index, letter := range theString {
		if unicode.IsNumber(letter) {
			fmt.Printf("at %d there is a letter %c\n", index, letter)

		}

	}

}
