package main

import (
	"fmt"
)

func main() {
	/* 
		if a variable does not have a type anotation 
		go allocate memory based on the archiquecture of the OS
	*/

	var iNumber int8 = 255
	var uNumber uint8 = -4
	var fNumber float32 = 5.7 // float64
	var str string = "I'm a string" // only use double quoted
	var aBool bool = false
	c := 10 + 8i 
}