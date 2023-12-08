package main

import (
	"fmt"
	"one"
	"two"
)

func SubMain() {
	fmt.Printf("SubMain\n")
	one.Print()
	two.Print()
}

func main() {
	SubMain()
}
