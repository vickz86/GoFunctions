package sliceint

import "fmt"

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
func AddToSliceUnder(theSlice []int, maxVal int) []int {

	//variable
	var intToAdd int
	var sliceToAdd []int
	var stringPrint = fmt.Sprintf("type int to add\n-1 to stop adding\nint must be <= %d", maxVal)

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

// AvergeSliceInt : get the average of the slice , return float with .1 precision
func AvergeSliceInt(theSlice []int) float64 {
	//variable
	lenSlice := float64(len(theSlice))
	var theSum int = 0

	//get the sum of int
	for _, nb := range theSlice {
		theSum += nb
	}

	SumFl := float64(theSum)

	ave := SumFl / lenSlice

	return ave

}
