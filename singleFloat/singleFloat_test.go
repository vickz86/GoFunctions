package singlefloat

import (
	"fmt"
	"testing"
)

func TestRoundSingleFloat(t *testing.T) {
	want := 12.6
	got := RoundSingleFloat(12.628, 1)

	if want != got {
		fmt.Printf("error RoundSingleFloat\n want = %v , got = %v", want, got)
	} else {
		fmt.Println("RoundSingleFloat has passed")
	}

}
