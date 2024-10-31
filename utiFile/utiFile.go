package utifile

import (
	"bufio"
	"fmt"
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


//write a slice of string to a file
func WriteFile(filepath string, strings []string) error {
	// Create or open the file
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close() // Ensure the file is closed when done

	writer := bufio.NewWriter(file) // Create a buffered writer

	// Write each string to the file with a newline
	for _, s := range strings {
		_, err := writer.WriteString(s + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}

	// Flush the buffered writer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("error flushing writer: %w", err)
	}

	return nil
}
