package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-bookshelf/cmd/author"
	"simple-bookshelf/test/cmd/mock"
)

type AuthorHandler struct {
	service author.Service
}

func NewAuthorHandler(service author.Service) *AuthorHandler {
	return &AuthorHandler{service}
}

func (h *AuthorHandler) GetAuthorById(context *gin.Context) {
	authorData := mock.GenerateAuthor("")
	context.JSON(http.StatusOK, authorData)
}
