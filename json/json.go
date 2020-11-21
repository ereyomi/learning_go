package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	pDetails := make(map[string]string)
	var name, address string
	fmt.Println("Enter your name")
	fmt.Scanf("%s", &name)
	fmt.Println("Enter your address")
	fmt.Scanf("%s", &address)
	pDetails["name"] = name
	pDetails["address"] = address

	fmt.Println("person Details [map]: ", pDetails)
	data, _ := json.Marshal(pDetails)
	fmt.Println("person Details [json]: ", string(data)) //or use Printf("%s", data)
}
