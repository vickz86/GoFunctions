package timeuti

import (
	"fmt"
	"testing"
)

func TestStringToSecond(t *testing.T) {
	strinGot,intGot,err := StringToSecond("4.45.34")
	if err!=nil{
		fmt.Println(err)
	}


	stringWant := "17334s"
	intWant := 17334

	if strinGot!=stringWant{
		fmt.Printf("want %v got %v",strinGot,stringWant)
	}
	
	if intGot!=intWant{
		fmt.Printf("want %v got %v",intGot,intWant)
	}
}