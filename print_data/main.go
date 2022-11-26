package main

import (
	"fmt"
)

func main () {
	hello_string := "Hello"
	world_string := "world"

	// use fmt.Println
	fmt.Println(hello_string, world_string)

	// use fmt.Printf
	name := "Holger"
	programming_lang := "Go"
	months_of_practicing := 2

	// use %s for strings, %d for integers and %v when data type is unknown
	fmt.Printf("%s, loves The %s Programming Language and has been practing it by %d months.\n", name, programming_lang, months_of_practicing)

	// use %T to know the data type
	fmt.Printf("Name contains a: %T and months_of_practicing an: %T.", name, months_of_practicing)
}