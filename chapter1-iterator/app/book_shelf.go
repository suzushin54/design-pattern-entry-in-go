package app

type BookShelf struct {
	books []*Book
	last  int
}

func NewBookShelf(maxsize int) *BookShelf {
	// BookShelfのサイズを指定
	books := make([]*Book, maxsize)
	return &BookShelf{books: books}
}

func (bs *BookShelf) GetBookAt(index int) *Book {
	return bs.books[index]
}

func (bs *BookShelf) AppendBook(book *Book) {
	bs.books[bs.last] = book
	bs.last++
}

func (bs *BookShelf) GetLength() int {
	return bs.last
}

// Iterator - BookShelf クラスに対応する. 本棚の本を繰り返し処理したい時に呼び出される
func (bs *BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(bs)
}
