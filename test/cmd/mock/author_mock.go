package mock

import (
	"simple-bookshelf/cmd/author"
)

func GenerateAuthor(name string) (result author.Author) {
	result = author.GenerateAuthor(name)
	return
}
