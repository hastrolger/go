package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]int)

	myMap["Jael"] = 29
	myMap["Wendy"] = 25

	// itarate map
	for k, v := range myMap {
		fmt.Println(k, v)
	}

	value, ok := myMap["Jael"]
	fmt.Println(value, ok)
}