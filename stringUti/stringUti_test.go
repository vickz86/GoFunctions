package stringuti

import (
	"fmt"
	"testing"
)

func TestSliceStringIndex(t *testing.T) {
	var stringTest string = "abc123"

	before := "abc1"
	after := "23"

	// get the value
	b, a, er := SliceStringIndex(stringTest, 3)

	//check error
	if er != nil{
		fmt.Println(er)
	}

	if before == b && after == a {
		fmt.Println("sliceStringIndex pass")
		} else {
			fmt.Printf("want %s %s and got %s %s ,sliceStringIndex error",before,after,b,a)
	}

}
