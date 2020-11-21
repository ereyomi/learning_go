package main

import "fmt"

func main() {

	var userInput float64
	fmt.Printf("Enter floating point number: ")
	fmt.Scanf("%g", &userInput)
	truncToInt := int(userInput)
	fmt.Printf("output %d ", truncToInt)

}
