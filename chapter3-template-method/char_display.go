package main

import (
	"fmt"
)

type CharDisplay struct {
	*abstractDisplay
	ch byte
}

func NewCharDisplay(ch byte) *CharDisplay {
	ad := &abstractDisplay{}
	return &CharDisplay{
		abstractDisplay: ad,
		ch:              ch,
	}
}

func (c *CharDisplay) Open() {
	fmt.Println("<<")
}

func (c *CharDisplay) Print() {
	fmt.Println(c.ch)
}

func (c *CharDisplay) Close() {
	fmt.Println(">>")
}
