package book

import "simple-bookshelf/cmd/author"

type Book struct {
	Id          string
	Isbn        string
	Title       string
	Authors     []author.Author
	Year        string
	Description string
	Pages       int
	Images      []string
}
