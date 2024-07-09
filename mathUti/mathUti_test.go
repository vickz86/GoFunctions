package mathuti

import (
	"fmt"
	"testing"
)

func TestPercentageVariation(t *testing.T) {
	want := 20.0
	got := PercentageVariation(50.0, 60.0)

	if want == got {
		fmt.Println("PercentageVariation has passed")
	} else {
		fmt.Printf("error PercentageVariation ,want %.1f and got %.1f\n", want, got)
	}

}

func TestPercVarWithLastAndOther(t *testing.T) {
	var sliceTOTest = []float64{5.0, 10.0, 15.0, 20.0}

	want := 300.0
	got, _ := PercVarWithLastAndOther(sliceTOTest, 0)

	if want != got {
		fmt.Printf("error PercVarWithLastAndOther, got %.1f and want %1.f\n", want, got)

	}

}
