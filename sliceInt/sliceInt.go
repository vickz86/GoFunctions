package sliceint

import "fmt"

// if len(sliceInt<fulldisplay) display all , else  display first ,last and middle
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
