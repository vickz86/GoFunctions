package timeuti

import (
	"errors"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
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


// change the year of a time.Time , 0 set to current year , other to add a year
func ChangeYear(oldDate time.Time, changeYear int) time.Time {
    // If changeYear is 0, set the year to the current year
    if changeYear == 0 {
        currentYear := time.Now().Year()
        yearDifference := currentYear - oldDate.Year()
        
        // Adjust the year of oldDate to the current year
        return oldDate.AddDate(yearDifference, 0, 0)
    } 

    // Otherwise, add or subtract years based on changeYear
    return oldDate.AddDate(changeYear, 0, 0)
}


// Convert a time.Duration to numbers of days , rounded to 1
func DurationToDays(theDur time.Duration) int{

    //Get the number of hours in theDur
    hours := theDur.Hours()

    // convert the hours to number of days
    nbDay := hours/24.0

    return int(math.Round(nbDay))
}
