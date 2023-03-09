package main

import "fmt"

func main9() {
	menu := map[string]float64{
		"soup": 4.99,
		"pie" : 7.99,
		"salad" : 6.99,
		"toffee pudding" : 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//loop thru a map
	for k, v := range menu { //k is the key +> key value pair
		fmt.Println(k, "-", v)
	}

	//ints as key type
	phonebook := map[int]string{
		234 : "Nigeria",
		456 : "Togo",
		789 : "Congo",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[234])

	phonebook[456] = "Algeria"
	phonebook[290] = "Benin Republic"

	for k, v := range phonebook { 
		fmt.Println(k, "-", v)
	}
}