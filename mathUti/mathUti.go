package mathuti

import "math"

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
