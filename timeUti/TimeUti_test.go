package timeuti

import (
	"fmt"
	"testing"
	"time"
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


func TestDurationToDayst (t *testing.T){

	//create the duration
	du,err := time.ParseDuration("75h30m")
	if err!= nil{
		fmt.Println(err)
	}

	got := DurationToDays(du)
	want := 3

	if got!=want{
		t.Errorf("want %v and got %v",want,got)
	}


}
