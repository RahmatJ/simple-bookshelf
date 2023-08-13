package mock

import (
	"simple-bookshelf/cmd/book"
)

func GenerateBook() (result book.Book) {
	result = book.GenerateBook("", "")
	return
}
