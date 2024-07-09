package slicefloat

import (
	"fmt"
	"testing"
)

func TestCompareLastToBefore(t *testing.T) {
	//var
	var testSlice = []float64{1.5, 2.5, 3.5, 5.0, 7.0}
	want1 := 2.0
	bool1, got1, _ := CompareLastToBefore(testSlice, 1)

	if want1 == got1 && bool1 == true {
		fmt.Println("CompareLastToBefore has passed")
	} else {
		fmt.Printf("error CompareLastToBefore ,got %.1f and want %.1f\n", got1, want1)
	}

}
