package main

import (
	"fmt"
	"strings"
)

func main8() {

	fmt.Println(getInitials("John Doe"))
}

func getInitials(n string) (string, string){ // return two strings
	str := strings.ToUpper(n)
	names := strings.Split(str, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], ""
}