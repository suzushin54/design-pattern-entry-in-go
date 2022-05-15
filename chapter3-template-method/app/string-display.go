package app

import "fmt"

// StringDisplay - ConcreteClass(具象クラスの役割)
// AbstractDisplayの抽象メソッドを全て実装する
type StringDisplay struct {
	*AbstractDisplay
	str   string
	width int
}

func NewStringDisplay(str string) *StringDisplay {
	ad := &AbstractDisplay{}
	sd := &StringDisplay{
		AbstractDisplay: ad,
		str:             str,
		width:           len(str),
	}
	sd.op = sd
	return sd
}

func (s *StringDisplay) Open() {
	s.printLine()
}

func (s *StringDisplay) Print() {
	fmt.Printf("|%s|\n", s.str)
}

func (s *StringDisplay) Close() {
	s.printLine()
}

func (s *StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < s.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
