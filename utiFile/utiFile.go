package utifile

import (
	"bufio"
	"log"
	"os"
)

// SliceStringFromFile : return a slice of string from a file name
func SliceStringFromFile(fileName string) ([]string, error) {
	//var to return
	var returnSlice []string

	//open file
	theFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer theFile.Close()

	//create a scanner
	newScanner := bufio.NewScanner(theFile)

	//for each line of the scanner
	for newScanner.Scan() {
		returnSlice = append(returnSlice, newScanner.Text())
	}

	if err := newScanner.Err(); err != nil {
		return nil, err
	}

	return returnSlice, nil

}
