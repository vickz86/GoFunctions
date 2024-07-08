package singlefloat

import (
	"math"
)

// RoundSingleFloat : round a float number based on decimal value
// if 0 than round to closest integer RoundSingleFloat(12.54,0) = 13
// eg : if bigger than 0 RoundSingleFloat(12.1234,1) = 12.1 , RoundSingleFloat(1.1234,3) = 1.123
func RoundSingleFloat(toRound float64, deciNb int) float64 {
	//variable

	newVal := toRound * math.Pow10(deciNb)
	// fmt.Println(newVal)
	rounded := math.Round(newVal)
	// fmt.Println(rounded)
	final := rounded / math.Pow10(deciNb)
	// fmt.Println(final)

	return final
}
