package main

import (
	"fmt"
	"math"
	"sort"
)

func main7() {
	
	Greeting("Lolade")
	sayBye("Kemi")
	fmt.Println("=====================================")
	var names  = []string{"Mark", "Mario", "Peach", "Bowl"};
	sort.Strings(names)
	cycleNames(names, Greeting)
	cycleNames(names, sayBye)
	fmt.Println("=====================================")
	fmt.Println("Area of Cirle")
	fmt.Println("=====================================")
	fmt.Printf("%0.3f", areaOfCircle(15))
}

func Greeting(name string){
	fmt.Printf("Good morning %v \n", name)
}

func sayBye(name string){
	fmt.Printf("Goodbye %v \n", name)
}

func cycleNames(n []string, f func(string)){
	for _, v := range n {  //_ means index, we don't need it
		f(v)
	}
}

func areaOfCircle(r float64) float64{ //return a float
	return math.Pi * r * r;
}