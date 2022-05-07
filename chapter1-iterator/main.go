package main

import (
	"fmt"
	"github.com/suzushin54/design-pattern-entry-in-go/chapter1-iterator/app"
	"os"
)

func main() {
	bookShelf := app.NewBookShelf(4)
	bookShelf.AppendBook(app.NewBook("Around the World in 80 days"))
	bookShelf.AppendBook(app.NewBook("Bible"))
	bookShelf.AppendBook(app.NewBook("Cinderella"))
	bookShelf.AppendBook(app.NewBook("Daddy-Long-Legs"))

	// 明示的に Iterator を使う方法
	it := bookShelf.Iterator()
	for it.HasNext() {
		book, err := it.Next()
		if err != nil {
			os.Exit(0)
		}
		fmt.Println(book.Name())
	}
}
