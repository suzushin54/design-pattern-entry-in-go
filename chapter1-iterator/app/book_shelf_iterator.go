package app

import (
	"fmt"
)

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func NewBookShelfIterator(bookShelf *BookShelf) *BookShelfIterator {
	return &BookShelfIterator{
		bookShelf: bookShelf,
		index:     0,
	}
}

func (bsi *BookShelfIterator) HasNext() bool {
	if bsi.index < bsi.bookShelf.GetLength() {
		return true
	}
	return false
}

// Next - 現在注目している本を返し、index を進める
// NOTE: 書籍では Java なので Iterator<Book> として実装している。
func (bsi *BookShelfIterator) Next() (*Book, error) {
	if !bsi.HasNext() {
		return nil, fmt.Errorf("no such element error")
	}

	book := bsi.bookShelf.GetBookAt(bsi.index)
	bsi.index++
	return book, nil
}
