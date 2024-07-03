package slicestring

import (
	"fmt"
	"testing"
)

var testString = []string{"victor", "amelie", "noemie", "jonathan"}
var nameSlice = "sliceOfName"

func TestCheckTypeInSlice(t *testing.T) {
	got := CheckTypeInSlice(testString, nameSlice)
	want := "amelie"

	if want != got {
		fmt.Printf("want: %s and got: %s", want, got)
	} else {
		fmt.Printf("CheckTypeInSlice has passed")
	}

}
