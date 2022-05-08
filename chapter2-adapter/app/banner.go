package app

import "fmt"

// Banner - 提供されているもの
type Banner struct {
	str string
}

func NewBanner(str string) *Banner {
	return &Banner{str: str}
}

func (b *Banner) ShowWithParen() {
	fmt.Printf("(%s)\n", b.str)
}

func (b *Banner) ShowWithAster() {
	fmt.Printf("*%s*\n", b.str)
}
