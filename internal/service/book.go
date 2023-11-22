package service

import (
	"github.com/asadbek21coder/demoproject/internal/entities"
	"github.com/asadbek21coder/demoproject/internal/repository"
)

type BooksService struct {
	repo repository.Books
}

func NewBooksService(repo repository.Books) *BooksService {
	return &BooksService{repo: repo}
}

func (s *BooksService) GetAllBooks(page, limit int) ([]*entities.Book, error) {
	return s.repo.GetAllBooks(page, limit)
}

func (s *BooksService) CreateBook(req *entities.CreateBookReq) (*entities.Book, error) {
	return s.repo.CreateBook(req)
}

func (s *BooksService) GetBookById(bookId int) (*entities.Book, error) {
	return s.repo.GetBookById(bookId)
}

func (s *BooksService) UpdateBook(id int, req *entities.UpdateBookReq) (*entities.Book, error) {
	return s.repo.UpdateBook(id, req)
}

func (s *BooksService) DeleteBook(id int) error {
	return s.repo.DeleteBook(id)
}
