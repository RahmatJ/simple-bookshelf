package author

import (
	"fmt"
	"simple-bookshelf/test/cmd/mock"
)

type Service interface {
	GetAuthors() ([]Author, error)
	GetAuthorById(id string) (Author, error)
	CreateAuthor(name string) (Author, error)
	PatchAuthorById(id string, name string) (Author, error)
	DeleteAuthorById(id string) error
}

//TODO(Rahmat): Please Fix circular dependencies, might need new package for it

type service struct {
}

func NewService() *service {
	return &service{}
}

func (s *service) GetAuthors() ([]Author, error) {
	authorData := mock.GenerateAuthor("")
	return []Author{authorData}, nil
}

func (s *service) GetAuthorById(id string) (Author, error) {
	authorData := mock.GenerateAuthor("")
	authorData.Id = id
	return authorData, nil
}

func (s *service) CreateAuthor(name string) (Author, error) {
	authorData := mock.GenerateAuthor(name)
	return authorData, nil
}
func (s *service) PatchAuthorById(id string, name string) (Author, error) {
	authorData := mock.GenerateAuthor(name)
	authorData.Id = id
	return authorData, nil
}

func (s *service) DeleteAuthorById(id string) error {
	fmt.Print(id)
	return nil
}
