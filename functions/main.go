package main

import (
	"fmt"
)

func addition(numA, numB int8) int8 {
	return numA + numB
}

func main() {
	result := addition(2,5)
	fmt.Println("Result of addition: ", result)
}