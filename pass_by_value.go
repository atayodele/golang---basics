package main

import "fmt"

func passValue() {

	//GROUP A types => strings, ints, bools, floats, arrays, structs

	name := "Tifa"
	name = updateName(name)
	fmt.Println(name)

	//GROUP B types => slices, maps, functions
	menu := map[string]float64{
		"pie" : 5.95,
		"ice cream" : 3.99,
	}
	updateMenu(menu)
	fmt.Println(menu)
}

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}