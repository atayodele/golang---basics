package main

import "fmt"

func pointer() {
	name := "Tifa"
	
	m := &name
	fmt.Println(name)
	updateNames(m)
	fmt.Println(name)
}

func updateNames(x *string) {
	*x = "wedge"
}