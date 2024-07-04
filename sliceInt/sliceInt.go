package sliceint

import (
	"fmt"
	"math"
	"slices"
)

// DisplaySliceInt: if len(sliceInt)<fulldisplay display all , else  display first ,last and middle
func DisplaySliceInt(sliceInt []int, fullDisplay int) {

	//varianle
	lenSl := len(sliceInt)

	//print whole
	if lenSl < fullDisplay {
		for index, inte := range sliceInt {
			fmt.Printf("%d : %d\n", index, inte)
		}

	} else /* if lenght slice is longer than fullDisplay */ {
		for index, inte := range sliceInt {
			if index == 0 || index == lenSl/2 || index == lenSl-1 {
				fmt.Printf("%d : %d\n", index, inte)
			}
		}
	}
}

// AddToSliceUnder: add int to the slice , value must be under maxVal , -1 to exit
// return a slice of int
func AddToSliceUnder(theSlice []int, maxVal int, addToWhat string) []int {

	//variable
	var intToAdd int
	var sliceToAdd []int
	var stringPrint = fmt.Sprintf("type int to add to %s\n-1 to stop adding\nint must be <= %d", addToWhat, maxVal)

	//add input int to sliceToAdd
	for {
		fmt.Println(stringPrint)
		_, err := fmt.Scan(&intToAdd)

		if err != nil {
			fmt.Print("error input, try again\n")
			continue
		}

		//if equal -1 break
		if intToAdd == -1 {
			break
		}

		if intToAdd < 0 || intToAdd > maxVal {
			fmt.Print("not in correct range!\n")
			continue
		}

		sliceToAdd = append(sliceToAdd, intToAdd)
		fmt.Printf("%d was added\n", intToAdd)

	}

	//add newly created slice to input slice
	theSlice = append(theSlice, sliceToAdd...)

	//return slice with appended new slice
	return theSlice
}

// CreateSliceFixNumber : create a slice of Int of fix number and return it
func CreateSliceFixNumber(fixNumber int) []int {
	var newSlice []int
	loop := 1
	var addNb int

	//loop and add
	for {
		fmt.Printf("type the %d number:", loop)
		fmt.Scan(&addNb)
		newSlice = append(newSlice, addNb)

		//check if reach fixNumber and return
		if loop == fixNumber {
			return newSlice
		}

		//increment counter
		loop++

	}

}

// AverageSliceInt : get the average of the slice , return float with .1 precision
func AverageSliceInt(theSlice []int) float64 {
	//variable
	lenSlice := float64(len(theSlice))
	var theSum int = 0

	//get the sum of int
	for _, nb := range theSlice {
		theSum += nb
	}

	SumFl := float64(theSum)

	ave := SumFl / lenSlice

	//convert result to .1 precision float
	ave *= 10
	ave = math.Round(ave)
	ave /= 10

	//return final value
	return ave
}

// GetNumberAbove : return a slice with all the number is slice above aboveNb
func GetNumberAbove(theSlice []int, aboveNb int) []int {
	var returnSlice []int

	// if number is above append to returnSlice
	for _, nb := range theSlice {
		if nb > aboveNb {
			returnSlice = append(returnSlice, nb)
		}

	}

	return returnSlice
}

// DeleteFirstOrLastSlice . delete last or first of slice of int : 0 for first , 1 for last
func DeleteFirstOrLastSlice(sliceInt []int, firstLast int) []int {
	//create a new slice which len = len(sliceInt)-1
	returnSlice := make([]int, len(sliceInt)-1)

	//define new slice
	var newSlice []int

	//create new slice slice
	if firstLast == 0 {
		//use slicing to remove first element of the slice
		newSlice = sliceInt[1:]
	} else if firstLast == 1 {
		//use slicing to remove last element of the slice
		newSlice = sliceInt[:len(sliceInt)-1]
	}

	//copy slice to returnslice
	copy(returnSlice, newSlice)

	//return copy slice
	return returnSlice

}

// RemoveStartEndSlice : remove element from slice int at either start or end of slice : 0 for start and 1 for end
func RemoveStartEndSlice(sliceInt []int, startEnd, nbtoRemove int) []int {
	//slice to return at the end, make sure has correct lenght
	sliceToReturn := make([]int, len(sliceInt)-nbtoRemove)
	// slice of slice
	var newSlice []int

	//remove at beginning
	if startEnd == 0 {
		newSlice = sliceInt[nbtoRemove:]
	} else /* remove at the end */ {
		newSlice = sliceInt[:len(sliceInt)-nbtoRemove]
	}

	//copy the result slice to a new slice
	copy(sliceToReturn, newSlice)

	return sliceToReturn

}

// ResizeSlice : resize sliceInt by remocing same numbers at beginning and end of the slice
// if number to remive is not even , it will remove more at beginning of the slice
func ResizeSlice(sliceInt []int, resizeSize int) []int {
	//Variable

	var isEven bool

	var temp1 []int
	var temp2 []int

	//get the total numbers to remove
	toRemove := len(sliceInt) - resizeSize
	// if negative , error
	if toRemove < 0 {
		fmt.Printf("error! your slice is lower than %d\n", resizeSize)
		return sliceInt
	}

	//create final return slice
	returnSlice := make([]int, resizeSize)

	//if to remove is even , remove same beginning and end
	if toRemove%2 == 0 {
		isEven = true
	}

	//create newslice
	if isEven {
		//removeAtStart
		temp1 = RemoveStartEndSlice(sliceInt, 0, toRemove/2)
		temp2 = RemoveStartEndSlice(temp1, 1, toRemove/2)

	} else {
		temp1 = RemoveStartEndSlice(sliceInt, 0, (toRemove/2)+1)
		//fmt.Println(temp1)
		temp2 = RemoveStartEndSlice(temp1, 1, toRemove/2)
	}

	//copy to a new slice
	copy(returnSlice, temp2)

	//return returnSlice
	return returnSlice

}

// KeepFirstLastSlice : keep first and last element of slice
func KeepFirstLastSlice(sliceInt []int) []int {
	var ReturnSlice []int

	//if len is 1 , return it
	if len(sliceInt) == 1 {
		fmt.Println("careful! one element in the slice")
		return sliceInt
	} else if len(sliceInt) == 2 /* if len is 2 , return it */ {
		return sliceInt
	} else /* loop append first and last */ {
		for index, element := range sliceInt {
			if index == 0 || index == (len(sliceInt)-1) {
				ReturnSlice = append(ReturnSlice, element)
			}

		}

	}
	return ReturnSlice

}

// SumSlice:  return sum of the slice
func SumSlice(sliceInt []int) int {
	var sum int

	for _, nb := range sliceInt {
		sum += nb
	}
	return sum

}

// IndexHigherVal : return highest valie and index (highestValue index)
// if multiple highest value , return first index of highest value
func HighestValueIndex(sliceint []int) (int, int) {
	//variable
	var highest int
	var index int

	// loop
	for index, intV := range sliceint {
		if index == 0 {
			highest = intV
		} else {
			if intV > highest {
				highest = intV

			}
		}
	}

	index = slices.Index(sliceint, highest)

	return highest, index

}
