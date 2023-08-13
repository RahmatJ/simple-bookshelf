package author

import (
	"fmt"
)

type Service interface {
	GetAuthors() ([]Author, error)
	GetAuthorById(id string) (Author, error)
	CreateAuthor(name string) (Author, error)
	PatchAuthorById(id string, name string) (Author, error)
	DeleteAuthorById(id string) error
}

type service struct {
}

func NewService() *service {
	return &service{}
}

func (s *service) GetAuthors() ([]Author, error) {
	authorData := GenerateAuthor("")
	return []Author{authorData}, nil
}

func (s *service) GetAuthorById(id string) (Author, error) {
	authorData := GenerateAuthor("")
	authorData.Id = id
	return authorData, nil
}

func (s *service) CreateAuthor(name string) (Author, error) {
	authorData := GenerateAuthor(name)
	return authorData, nil
}
func (s *service) PatchAuthorById(id string, name string) (Author, error) {
	authorData := GenerateAuthor(name)
	authorData.Id = id
	return authorData, nil
}

func (s *service) DeleteAuthorById(id string) error {
	fmt.Print(id)
	return nil
}
