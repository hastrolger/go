// a way to organize goroutines
package main

import (
	"fmt"
)

func say(text string, c chan string) {
	c <- text
}

func main() {
	// make a channel and define 
	// what data type is gonna be used in it
	// how much data is gonna process
	c := make(chan string, 1) 

	fmt.Println("Hello")

	go say("Bye", c)

	fmt.Println(<- c)

	close(c)

	// iterate a chan and selects its elements 
	cInt := make(chan int, 2)
	cInt <- 22
	cInt <- 11

	//fmt.Println(len(cInt), cap(cInt))

	for i := range cInt {
		fmt.Println(i)
	}
}