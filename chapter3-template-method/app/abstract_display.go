package app

type Operation interface {
	Open()
	Print()
	Close()
}

// AbstractDisplay - AbstractClass(抽象クラスの役割)
type AbstractDisplay struct {
	op Operation
}

// Display - AbstractDisplayが持つ3つのメソッドを使う、テンプレートメソッド
func (a *AbstractDisplay) Display() {
	a.op.Open()
	for i := 0; i < 5; i++ {
		a.op.Print()
	}
	a.op.Close()
}
