package main

import "fmt"

func main1() {
	//STRING
	var name1 string = "Mario"
	var name2 = "Mark"
	var name3 string
	fmt.Println(name1, name2, name3);

	name3 = "peach"
	name2 = "boucer"
	fmt.Println(name1, name2, name3);

	name4 := "Yagi level" // this can be done inside a function not outside a function
	
	fmt.Println(name4);

	//INTERGERS

	var age1 int = 20
	var age2 = 30
	age3 := 40
	fmt.Println(age1, age2, age3);

	//bits & memory
	var num1 int8 = 45
	var num2 int8 = -128
	var num3 uint8 = 255
	fmt.Println(num1, num2, num3);

	//float
	var score1 float32 = 25.98
	var score2 float64 = 88375923759436384657348.7
	score3 := 453.44

	fmt.Println(score1, score2, score3);
}