package mock

import (
	"simple-bookshelf/cmd/author"
	"simple-bookshelf/cmd/book"
)

func GenerateBook() (result book.Book) {
	authorData := GenerateAuthor()
	result = book.Book{
		Id:    "1",
		Isbn:  "1000",
		Title: "Detective Conan",
		Pages: 50, Year: "1995",
		Authors:     []author.Author{authorData},
		Description: "Dummy",
		Images:      []string{"cover.png"},
	}
	return
}
