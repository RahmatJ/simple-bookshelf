package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-bookshelf/cmd/author"
	"simple-bookshelf/cmd/book"
)

type BookHandler struct {
}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

// GetBook Get All Book
func (h *BookHandler) GetBook(context *gin.Context) {
	authorData := author.Author{Id: "1", Name: "Aoyama Gosho"}
	authors := []author.Author{authorData}
	bookData := book.Book{
		Id:    "1",
		Isbn:  "1000",
		Title: "Detective Conan",
		Pages: 50, Year: "1995",
		Authors:     authors,
		Description: "Dummy",
		Images:      []string{"cover.png"},
	}

	context.JSON(http.StatusOK, bookData)
}

// CreateBook Create Book
func (h *BookHandler) CreateBook(context *gin.Context) {
	context.JSON(http.StatusCreated, nil)
}

// UpdateBook Update Book
func (h *BookHandler) UpdateBook(context *gin.Context) {
	authorData := author.Author{Id: "1", Name: "Aoyama Gosho"}
	authors := []author.Author{authorData}
	bookData := book.Book{
		Id:    "1",
		Isbn:  "1000",
		Title: "Detective Conan",
		Pages: 50, Year: "1995",
		Authors:     authors,
		Description: "Dummy",
		Images:      []string{"cover.png"},
	}
	context.JSON(http.StatusOK, bookData)
}

// DeleteBook Delete Book
func (h *BookHandler) DeleteBook(context *gin.Context) {
	context.JSON(http.StatusNoContent, nil)
}
