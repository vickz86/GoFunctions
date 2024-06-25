package sliceint

import (
	"fmt"
	"slices"
	"testing"
)

func TestAvergeSliceInt(t *testing.T) {
	testSlice := []int{3, 5, 3}
	whatGot := AverageSliceInt(testSlice)
	whatWant := 3.7
	if whatGot == whatWant {
		fmt.Println("succes of test")
	} else {
		fmt.Printf("error: got %f , wanted %f", whatGot, whatWant)
	}

}

func TestGetNumberAbove(t *testing.T) {
	testSlice := []int{6, 12, 8, 9, 14, 18}

	got := GetNumberAbove(testSlice, 10)
	var want = []int{12, 14, 18}
	isEqual := slices.Equal(got, want)
	if isEqual {
		fmt.Println("succes of test")
	} else {
		fmt.Printf("error: got %v , wanted %v", got, want)
	}

}
