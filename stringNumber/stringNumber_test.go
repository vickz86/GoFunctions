package stringnumber

import (
	"fmt"
	"slices"
	"testing"
)

func TestSliceIndexConnectedNumbers(t *testing.T) {
	// var theString = "hello128baby545mama"
	var theString2 = "ba987cd456"

	want := []int{2, 3, 4}
	got := SliceIndexConnectedNumbers(theString2)

	if slices.Equal(want, got) {
		fmt.Println("working!")
	} else {
		fmt.Printf("error!want:%v and got %v", want, got)
	}

	SliceIndexConnectedNumbers(theString2)

}
