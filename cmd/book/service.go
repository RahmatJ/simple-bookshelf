package book

type Service interface {
	GetBooks() ([]Book, error)
	GetBookById(id string) (Book, error)
	CreateBook(title string) (Book, error)
	PatchBookById(id string, title string) (Book, error)
	DeleteBookById(id string) (Book, error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}

func (s *service) GetBooks() ([]Book, error) {
	book := GenerateBook("", "")
	return []Book{book}, nil
}

func (s *service) GetBookById(id string) (Book, error) {
	book := GenerateBook(id, "")
	return book, nil
}

func (s *service) CreateBook(title string) (Book, error) {
	book := GenerateBook("", title)
	return book, nil
}

func (s *service) PatchBookById(id string, title string) (Book, error) {
	book := GenerateBook(id, title)
	return book, nil
}

func (s *service) DeleteBookById(id string) (Book, error) {
	book := GenerateBook(id, "")
	return book, nil
}
