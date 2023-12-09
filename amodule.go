package main

import (
	"fmt"

	"amodule/internal/two"
	"amodule/one"
)

func SubMain() {
	fmt.Printf("SubMain\n")
	one.Print()
	two.Print()
}

func main() {
	SubMain()
}
