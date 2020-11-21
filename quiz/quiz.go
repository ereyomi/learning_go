package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	readFile, err := os.Open("file.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
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
