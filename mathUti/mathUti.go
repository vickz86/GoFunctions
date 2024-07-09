package mathuti

import (
	"errors"
	"math"
)

// PercentageVariation : get the percentage variation between 2 float
func PercentageVariation(float1, float2 float64) float64 {
	//first substract
	sub := float2 - float1

	//div
	div := sub / math.Abs(float1)

	//mul by 100
	final := div * 100

	return final

}

// PercVarWithLastAndOther : get percentage variation between last value of slice of float and other
// 0 to compare with first of slice , 1 to walk backward
// return percentage variation
func PercVarWithLastAndOther(slicefloat []float64, ToCompareWith int) (float64, error) {
	// check for error
	if ToCompareWith > len(slicefloat) {
		return 0.0, errors.New("error,CompareWith int is too big")
	}

	//get last value
	lastVal := slicefloat[len(slicefloat)-1]

	//get value to compare with
	var CompVal float64
	if ToCompareWith == 0 {
		CompVal = slicefloat[0]
	} else {
		CompVal = slicefloat[len(slicefloat)-1-ToCompareWith]
	}

	//Create the variation using the PercentageVariation function
	difVal := PercentageVariation(CompVal, lastVal)

	//return the value
	return difVal, nil

}
