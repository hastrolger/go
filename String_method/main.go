package main

import (
	"fmt"
)

type pc struct {
	ram int 
	brand string
	disk int
}

func (myPc pc) String() {
	fmt.Printf("PC capabilities: brand %s, %d of ram and %d of disk.", myPc.brand, myPc.ram, myPc.disk)
}

func main() {
	myPc := pc {ram: 4, brand: "lenovo", disk: 2}
	myPc.String()
}