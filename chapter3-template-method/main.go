package main

import (
	"fmt"
	"github.com/suzushin54/design-pattern-entry-in-go/chapter3-template-method/app"
)

func main() {
	d1 := app.NewCharDisplay('H')
	d1.Display()

	fmt.Println()

	d2 := app.NewStringDisplay("Hello World.")
	d2.Display()
}
