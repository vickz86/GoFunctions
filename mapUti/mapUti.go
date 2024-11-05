package maputi

import "fmt"

//declare a typr
type WordChooser map[string]string

// choose a word from choose and return it
func ChooseWordFromMap(todo string,choose WordChooser)string{
	// define return string
	var returnString string

	
	//print the string
	fmt.Println(todo)

	//print all element from choose
	for key,value := range choose{
		fmt.Printf("%s : %s\n",key,value)
	}

	//get element
	fmt.Scan(&returnString)

	//chek if returnstring is in the map
	v,ok := choose[returnString]
	if !ok{
		fmt.Printf("%v is not define!",v)

	}

	
	
	return ""

}