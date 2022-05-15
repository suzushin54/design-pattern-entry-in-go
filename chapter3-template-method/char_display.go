package main

import (
	"fmt"
)

// CharDisplay - AbstractDisplayの抽象メソッドを全て実装するサブクラスの役割
type CharDisplay struct {
	*AbstractDisplay
	ch byte
}

func NewCharDisplay(ch byte) *CharDisplay {
	ad := &AbstractDisplay{}
	cd := &CharDisplay{
		AbstractDisplay: ad,
		ch:              ch,
	}

	// I/Fを満たすので抽象メソッドを実装したCharDisplayを持たせられる
	cd.op = cd
	return cd
}

func (c *CharDisplay) Open() {
	fmt.Print("<<")
}

func (c *CharDisplay) Print() {
	fmt.Print(string(c.ch))
}

func (c *CharDisplay) Close() {
	fmt.Print(">>")
}
