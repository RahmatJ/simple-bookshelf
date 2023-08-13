package book

import "simple-bookshelf/cmd/author"

func GenerateBook(id string, title string) (result Book) {
	if id == "" {
		id = "1"
	}
	if title == "" {
		title = "Detective Conan"
	}
	authorData := author.GenerateAuthor("")
	result = Book{
		Id:    id,
		Isbn:  "1000",
		Title: title,
		Pages: 50, Year: "1995",
		Authors:     []author.Author{authorData},
		Description: "Dummy",
		Images:      []string{"cover.png"},
	}
	return
}
