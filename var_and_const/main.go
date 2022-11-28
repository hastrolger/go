package main

import (
	"fmt"
)

func main() {
	// constants
	// it's not possible to infer a value type using ":=" operator
	const PI = 3.1415

	fmt.Printf("The value of PI is: %v and its type is: %T\n", PI, PI)

	// variables 
	// if any variable is unused the program cannot run
	/*
	inferredVar := false
	var typeGiven int = 2
	var justInitialized string 
	*/

	// square area
	const baseCuadrado = 10
	squareArea := baseCuadrado * baseCuadrado
	fmt.Println("Area of a square: ", squareArea)

}