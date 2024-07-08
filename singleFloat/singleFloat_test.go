package singlefloat

import (
	"fmt"
	"testing"
)

func TestRoundSingleFloat(t *testing.T) {
	want := 12.628
	got := RoundSingleFloat(12.6284, 3)

	if want != got {
		fmt.Printf("error RoundSingleFloat\n want = %v , got = %v", want, got)
	} else {
		fmt.Println("RoundSingleFloat has passed")
	}

}
