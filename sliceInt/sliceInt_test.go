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

func TestDeleteFirstOrLastSlice(t *testing.T) {
	var testSlice = []int{2, 4, 6, 8}
	want := []int{2, 4, 6}
	got := DeleteFirstOrLastSlice(testSlice, 1)
	//check if equal
	isEqual := slices.Equal(want, got)
	//print result
	if isEqual {
		fmt.Println("test has passed")
	} else {
		fmt.Printf("got:%v and wanted %v", got, want)
	}

}

func TestRemoveStartEndSlice(t *testing.T) {
	var testSliceInt = []int{2, 4, 6, 8, 10, 12, 14}
	want := []int{2, 4, 6, 8}
	got := RemoveStartEndSlice(testSliceInt, 1, 3)
	isEqual := slices.Equal(want, got)
	//print result
	if isEqual {
		fmt.Println("test has passed")
	} else {
		fmt.Printf("got:%v and wanted %v", got, want)
	}

}

func TestResizeSlice(t *testing.T) {
	var testSliceInt = []int{2, 4, 6, 8, 10, 12, 14, 16}
	want := []int{6, 8, 10, 12, 14}
	got := ResizeSlice(testSliceInt, 5)
	isEqual := slices.Equal(want, got)
	//print result
	if isEqual {
		fmt.Println("test has passed")
	} else {
		fmt.Printf("error for test ResizeSliceInt\n")
		fmt.Printf("got:%v and wanted %v", got, want)
	}

}

func TestKeepFirstLastSlice(t *testing.T) {
	var test = []int{2, 3, 4, 5, 6}
	got := KeepFirstLastSlice(test)
	want := []int{2, 6}

	if slices.Equal(got, want) {
		fmt.Println("KeepFirstLastSlice has passed")
	} else {
		fmt.Printf("want %v , got %v", want, got)
	}

}

func TestSumSlice(t *testing.T) {
	var test = []int{2, 3, 4}
	got := SumSlice(test)
	want := 9

	if got == want {
		fmt.Println("SumSlice has passed")
	} else {
		fmt.Printf("error ,want %v , got %v\n", want, got)
	}

}

func TestHighestValueIndex(t *testing.T) {
	var test = []int{2, 3, 4, 4, 8, 5, 2, 12, 3}
	gotVal, gotIn := HighestValueIndex(test)
	wantVal, wantIn := 12, 7

	if gotVal == wantVal && gotIn == wantIn {
		fmt.Println("HighestValueIndex has passed")
	} else {
		fmt.Printf("error HighestValueIndex want %v , got %v and want %v , got %v", wantVal, gotVal, wantIn, gotIn)
	}

}
