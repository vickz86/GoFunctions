package sliceint

import (
	"fmt"
	"testing"
)

func TestAvergeSliceInt(t *testing.T) {
	testSlice := []int{2, 2}
	whatGot := AvergeSliceInt(testSlice)
	whatWant := 2.0
	if whatGot == whatWant {
		fmt.Println("succes of test")
	} else {
		fmt.Printf("error: got %f , wanted %f", whatGot, whatWant)
	}

}
