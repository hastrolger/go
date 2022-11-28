package main

import (
	"fmt"
)

func main() {
	// defer
	defer fmt.Println("I'll be the last sentence to exec")
	fmt.Println("I'll go first")

	// continue and break
	for i := 0; i <= 4; i++ {
		if i == 2 {
			continue
		} else if i == 3 {
			break
		}
		fmt.Println(i)
	}
}