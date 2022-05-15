package app

import (
	"fmt"
)

// CharDisplay - ConcreteClass(具象クラスの役割)
// AbstractDisplayの抽象メソッドを全て実装する
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
