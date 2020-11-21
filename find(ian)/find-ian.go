package main

import (
	"fmt"
	"strings"
)

func main() {

	var userInput string
	fmt.Print("Enter Sting: ")
	fmt.Scanf("%s", &userInput)
	convertToLower := strings.ToLower(userInput)

	checkA := strings.HasPrefix(convertToLower, "i")
	checkB := strings.HasSuffix(convertToLower, "n")
	checkC := strings.ContainsAny(convertToLower, "a")
	if checkA && checkB && checkC {
		println("Found")
	} else {
		println("not found")
	}

}
