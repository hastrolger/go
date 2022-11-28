package main

import (
	"fmt"
)

type car struct {
	brand string
	year int16
}

func main() {
	myCar := car{
		brand: "Ford",
		year: 2009,
	}

	fmt.Println(myCar)

	var otherCar car
	otherCar.brand = "KIA"
	fmt.Println(otherCar)
}