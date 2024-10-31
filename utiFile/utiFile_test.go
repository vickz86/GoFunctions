package utifile

import "testing"

func TestWriteFile(t *testing.T) {
	testSlice := []string{
		"Hello World",
		"Go Language",
		"Write File",
		"Read File",
		"Test Function",
		"Slice Example",
	}

	WriteFile("data.txt",testSlice)

}