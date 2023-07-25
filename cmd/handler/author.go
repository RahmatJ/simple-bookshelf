package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-bookshelf/test/cmd/mock"
)

type AuthorHandler struct {
}

func NewAuthorHandler() *AuthorHandler {
	return &AuthorHandler{}
}

func (h *AuthorHandler) GetAuthorById(context *gin.Context) {
	authorData := mock.GenerateAuthor()
	context.JSON(http.StatusOK, authorData)
}
