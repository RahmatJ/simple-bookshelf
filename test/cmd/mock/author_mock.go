package mock

import "simple-bookshelf/cmd/author"

func GenerateAuthor() (result author.Author) {
	result = author.Author{}
	result.Id = "1"
	result.Name = "Aoyama Gosho"
	return
}
