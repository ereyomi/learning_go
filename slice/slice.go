package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	userArray := make([]int, 3)

	i := 0

	fmt.Println("You can choose to Exit the program by inputting X to exit")

	for {

		if i < 3 {
			fmt.Print("your input: ")
		} else {
			fmt.Println("Your slice have reach maximum cap. Input X to exit")
		}
		var userInputData string

		fmt.Scan(&userInputData)

		if strings.ToLower(userInputData) == "x" {
			break
		}
		uData, err := strconv.Atoi(userInputData)
		if i < 3 {
			if err != nil {
				fmt.Println("You inputed a wrong data")
				continue
			} else {
				userArray[i] = uData
				i++
				fmt.Print("appended to array\n")

			}
		}

	}

	fmt.Println("unsort array: ", userArray)
	sort.Ints(userArray[:])
	fmt.Println("sorted array: ", userArray)

}
