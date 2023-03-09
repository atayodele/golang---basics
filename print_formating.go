package main

import "fmt"

func main2() {
	//Printing and Formating string

	fmt.Print("Hello,")
	fmt.Println("Hello, ninjas!")

	age := 35
	name := "Timi"
	fmt.Println("My age is", age, "and my name is", name)

	//Formatting
	//printf() - formatted string
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %v and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you score %f points! \n", 255.55)
	fmt.Printf("you score %0.1f points! \n", 255.55)

	//Sprintf - save formatted strings
	var str = fmt.Sprintf("my age is %v and my name is %q \n", age, name)
	fmt.Println("The saved string is:", str)

}