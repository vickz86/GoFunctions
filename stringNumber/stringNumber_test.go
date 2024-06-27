package stringnumber

import (
	"fmt"
	"slices"
	"testing"
)

func TestSliceIndexConnectedNumbers(t *testing.T) {
	// var theString = "hello128baby545mama"
	var theString2 = "ba9876cd456"

	want := []int{2, 5}
	got := SliceIndexConnectedNumbers(theString2)

	if slices.Equal(want, got) {
		fmt.Println("working!")
	} else {
		fmt.Printf("error!want:%v and got %v", want, got)
	}

}

func TestGetStringBeforeAfterIndex(t *testing.T) {
	var testString = "test55phrase"

	want := "test"
	got := GetStringBeforeAfterIndex(testString, 4, 0)
	if want == got {
		fmt.Println("working!")
	} else {
		fmt.Printf("error!want:%v and got %v", want, got)
	}

}

func TestGetStringBetweenslice(t *testing.T) {
	// var testString = "test555phrase"
	// var theSlice = []int{4, 6}

	// var testString2 = "mama1ma"
	// var theSlice2 = []int{4}

	var testString3 = "ubba00100chaka"
	var theSlice3 = []int{4, 8}

	want := "00100"
	got := GetStringBetweenslice(testString3, theSlice3)
	if want == got {
		fmt.Println("working!")
	} else {
		fmt.Printf("error!want:%v and got %v", want, got)
	}

}

func TestStringToInt(t *testing.T) {
	want := 12

	got := StringToInt("10", 2)

	if want == got {
		fmt.Println("StringToInt has passed")
	} else {
		fmt.Printf("error StringToInt , got %v ,want %v", got, want)
	}
}

func TestAddToIntStringInString(t *testing.T) {
	var testString = "test123in45"
	want := "test126in45"

	got := AddToIntStringInString(testString, 3)

	if want == got {
		fmt.Println("AddToIntStringInString has passed")
	} else {
		fmt.Printf("error StringToInt , got %v ,want %v", got, want)
	}
}
