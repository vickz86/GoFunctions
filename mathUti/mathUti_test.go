package mathuti

import (
	"fmt"
	"testing"
)

func TestPercentageVariation(t *testing.T) {
	want := 10.0
	got := PercentageVariation(50.0, 60.0)

	if want == got {
		fmt.Println("PercentageVariation has passed")
	} else {
		fmt.Printf("error PercentageVariation ,want %.1f and got %.1f", want, got)
	}

}
