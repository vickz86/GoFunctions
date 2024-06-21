package sliceint

import "fmt"

func DisplaySliceInt(sliceInt []int, fullDisplay int) {
	//if len of sliceInt<fulldisplay display all , otherwisw display first ,last and middle

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
