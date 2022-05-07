package app

type Book struct {
	name string
}

func NewBook(name string) *Book {
	return &Book{name: name}
}

func (b Book) Name() string {
	return b.name
}
