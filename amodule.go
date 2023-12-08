package amodule

import (
	"fmt"

	"github.com/tildeleb/amodule/one"
	"github.com/tildeleb/amodule/two"
)

func SubMain() {
	fmt.Printf("SubMain\n")
	one.Print()
	two.Print()
}

/*
func main() {
	SubMain()
}
*/
