package main

import (
	"fmt"
	"sort"
	"strings"
)

func main4() {

	greeting := "hello there friends!"
	fmt.Println(strings.Contains(greeting, "ello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "Hi"))
	fmt.Println(strings.Index(greeting, "th"))
	fmt.Println(strings.Split(greeting, "there"))
	fmt.Println(strings.ToLower(greeting))
	fmt.Println(strings.ToUpper(greeting))

	ages := []int{45, 20,35,75,60, 50, 25};
	sort.Ints(ages)
	fmt.Println(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)
	
	names := []string{"Mark", "Mario", "Peach", "Bowl"};
	sort.Strings(names) // in alphabetical order
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "Bowl"))
}