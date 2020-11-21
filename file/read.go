package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {

	var filename string

	fmt.Println("Enter File name (e.g test.txt): ")
	_, err := fmt.Scan(&filename)
	readFile, fileError := os.Open(filename)

	if err != nil || fileError != nil {
		fmt.Println("An error occurred!, couldnt get the file")
		return
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	//sliceLen := len(fileTextLines)

	var mySliceStruct []Person

	for _, eachline := range fileTextLines {

		words := strings.Split(eachline, " ")
		// fmt.Println(len(words))
		wordslenght := len(words)

		if wordslenght == 2 {
			mySliceStruct = append(mySliceStruct, Person{words[0], words[1]})
		} else {

			fmt.Println("Wrong format of string!")
		}

	}

	//fmt.Print(mySliceStruct)

	for _, fullName := range mySliceStruct {
		fmt.Printf("%s %s\n", fullName.fname, fullName.lname)
	}
}
