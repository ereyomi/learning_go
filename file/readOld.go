/*package main

import (
	"fmt"
	"log"
	"os"
)
//Seek
func main() {
	data := make([]byte, 100)
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		//d, _ := file.Read(data)
		file.Read(data)
		fmt.Print(string(data))
		file.Close()
	}

}
*/