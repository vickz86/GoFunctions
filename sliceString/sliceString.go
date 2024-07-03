package slicestring

import (
	"fmt"
	"slices"
)

// CheckTypeInSlice : check what you type is in the slice of string ,it will return it.
// toPrint parameter will print description of the slice
func CheckTypeInSlice(sliceToCheck []string, toPrint string) string {
	var inputString string

	for {
		// get the input
		fmt.Printf("type string you want to check is in %s\nq to exit\na print all\n", toPrint)
		fmt.Scan(&inputString)

		// exit is type q
		if inputString == "q" {
			return ""
		}

		//print all
		if inputString == "a" {
			for _, stringSl := range sliceToCheck {
				fmt.Println(stringSl, " is present")

			}
		}

		//check if present or not
		if slices.Contains(sliceToCheck, inputString) {
			return inputString
		} else {
			fmt.Println("coudnt find ", inputString, " type a new one")
		}

	}

}
