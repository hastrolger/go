package main

import (
	"fmt"
	"sync"
)

func say(msg string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hello")
	wg.Add(1)

	go say("world", &wg)
}