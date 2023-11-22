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
