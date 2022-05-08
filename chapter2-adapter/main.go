package main

import (
	"github.com/suzushin54/design-pattern-entry-in-go/chapter2-adapter/app"
)

func main() {
	p := app.NewPrintBanner("Hello")
	p.PrintWeak()
	p.PrintStrong()
}
