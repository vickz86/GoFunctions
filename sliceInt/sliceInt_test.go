package sliceint

import (
	"fmt"
	"testing"
)

func TestAvergeSliceInt(t *testing.T) {
	testSlice := []int{3, 5, 3}
	whatGot := AvergeSliceInt(testSlice)
	whatWant := 3.7
	if whatGot == whatWant {
		fmt.Println("succes of test")
	} else {
		fmt.Printf("error: got %f , wanted %f", whatGot, whatWant)
	}

}
