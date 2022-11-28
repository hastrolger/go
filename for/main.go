package main

import(
	"fmt"
)

func main() {
	// for conditional
	for i := 0; i >= 5;  i++ {
		fmt.Println("Iteration number: ", i)
	}

	// for while
	counter := 0
	for counter < 5 {
		fmt.Println("Counter: ", counter)
		counter++
	}

	// for forever
	for {
		fmt.Println("You can only stop me by pressing ctrl + c...")
	}
}