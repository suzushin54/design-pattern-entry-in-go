package app

// PrintBanner - アダプターの役割を担う。変換装置。
type PrintBanner struct {
	// NOTE: 書籍では Banner クラスを継承
	Banner
}

func NewPrintBanner(str string) *PrintBanner {
	banner := NewBanner(str)
	return &PrintBanner{Banner: *banner}
}

// NOTE: 下記のメソッドは PrintBanner が勝手に作っているのではなく、
//       Print interface で必要なものとして定義しているもの。

func (pb *PrintBanner) PrintWeak() {
	pb.ShowWithParen()
}

func (pb *PrintBanner) PrintStrong() {
	pb.ShowWithAster()
}
