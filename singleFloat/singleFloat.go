package singlefloat

import (
	"fmt"
	"math"
)

// RoundSingleFloat : round a float number based on decimal value
// if 0 than round to closest integer RoundSingleFloat(12.54,0) = 13
// eg : if bigger than 0 RoundSingleFloat(12.1234,1) = 12.1 , RoundSingleFloat(1.1234,3) = 1.123
func RoundSingleFloat(toRound float64, deciNb int) float64 {
	//variable

	switch deciNb {
	case 0:
		return math.Round(toRound)
	case 1:
		toRound *= 10
		toRound = math.Round(toRound)
		return toRound / 10
	default:
		fmt.Println("error , decimal number not correct")
		return 0.0
	}

}
