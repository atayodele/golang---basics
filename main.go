package main

import "fmt"

func main() {

	sayHello("Ayo")

	for _, v := range points{
		fmt.Println(v)
	}

	// FOR STRUCT
	myBill := createBill()
	promptOptions(myBill)

	// INTERFACES
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
}
