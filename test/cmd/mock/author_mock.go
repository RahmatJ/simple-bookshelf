package mock

import "simple-bookshelf/cmd/author"

func GenerateAuthor(name string) (result author.Author) {
	if name == "" {
		name = "Aoyama Gosho"
	}
	result = author.Author{}
	result.Id = "1"
	result.Name = name
	return
}
