package stringuti

import (
	"testing"
)

func TestSliceStringIndex(t *testing.T) {
	tests := []struct {
		input    string
		index    int
		expectedBefore string
		expectedAfter  string
		expectError    bool
	}{
		{"abc123", 3, "abc1", "23", false},
		{"abc123", 0, "a", "bc123", false},
		{"abc123", 5, "abc123", "", false},
		{"abc123", 6, "", "", true}, // Out of range
		{"abc123", -1, "", "", true}, // Negative index
	}

	for _, tt := range tests {
		b, a, err := SliceStringIndex(tt.input, tt.index)

		if (err != nil) != tt.expectError {
			t.Errorf("SliceStringIndex(%q, %d) returned error: %v, want error: %v", tt.input, tt.index, err, tt.expectError)
		}
		if b != tt.expectedBefore || a != tt.expectedAfter {
			t.Errorf("SliceStringIndex(%q, %d) = %q, %q; want %q, %q", tt.input, tt.index, b, a, tt.expectedBefore, tt.expectedAfter)
		}
	}
}


func TestSliceStringByte(t *testing.T){
	var theString string = "darla you are mine forever"
	var theLetter string = "b"

	SliceStringString(theString,theLetter)

}
