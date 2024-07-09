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

func TestCompare2Float(t *testing.T) {
	float1 := 4.5
	float2 := 2.0
	want := 2.5
	got := Compare2Float(float1, float2)

	if want != got {
		fmt.Printf("error Compare2Float\n want = %v , got = %v", want, got)
	} else {
		fmt.Println("Compare2Float has passed")
	}

}
