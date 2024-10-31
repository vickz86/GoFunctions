package timeuti

import (
	"errors"
	"math"
	"slices"
	"strconv"
	"strings"
)

// convert an input of type hh.mm.ss into either "xxxs" or an int xxx
// return "xxxs" (string) , xxx (int)
func StringToSecond(timeString string) (string, int,error) {
    var nbDots int
    var totalSecond int

    //count the number of dot
    nbDots = strings.Count(timeString,".")
    //if more than 2 dots ,error
    if nbDots>2{
        return "",0,errors.New("ERROR,too much . in your input")
    }

    //split the timestring based on dots
    elements := strings.Split(timeString,".")
    slices.Reverse(elements)

    //loop through all elements
    for nb,el := range elements{

        //try to covert each element to a float
        floatTime,er := strconv.ParseFloat(el,64)
        if er!= nil{
            return "",0,errors.New("ERROR,an element in input is not an int")
        }

        //create the multiplicator based on power
        val := math.Pow(60.0,float64(nb))

        secondVal := floatTime*val
        totalSecond += int(secondVal)
    }

    //convert the int into a string
    outString := strconv.Itoa(totalSecond)+"s"
    
    return outString,totalSecond,nil
}