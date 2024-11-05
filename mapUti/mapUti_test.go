package maputi

import "testing"

func TestChooseWordFromMap(t *testing.T) {
	
	testMap := WordChooser{
		"add":    "add element",
		"search": "search for specific elements",
		"delete": "delete specific elements",
	}

	ChooseWordFromMap("what to do",testMap)

}