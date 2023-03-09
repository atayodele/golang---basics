package main

import (
	"fmt"
	"sort"
)

func main5() {

	for i := 0; i < 5; i++ {
		fmt.Println("Value of x is:", i)
	}

	names := []string{"Mark", "Mario", "Peach", "Bowl"};
	sort.Strings(names)
	for x := 0; x < len(names); x++ {
		fmt.Println(names[x])
	}

	for index, value := range names {
		fmt.Printf("The position at index %v is %v \n", index, value)
	}
	// fmt.Println("Hello, ninjas!")
}