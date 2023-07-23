package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-bookshelf/cmd/author"
)

type AuthorHandler struct {
}

func NewAuthorHandler() *AuthorHandler {
	return &AuthorHandler{}
}

func (h *AuthorHandler) GetAuthorById(context *gin.Context) {
	authorData := author.Author{Id: "1", Name: "Aoyama Gosho"}
	context.JSON(http.StatusOK, authorData)
}
