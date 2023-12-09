package amodule

import (
	"fmt"

	"github.com/tildeleb/amodule/internal/two"
	"github.com/tildeleb/amodule/one"
)

func SubMain() {
	fmt.Printf("SubMain\n")
	one.Print()
	two.Print()
}
