package service

import (
	"github.com/asadbek21coder/demoproject/internal/entities"
	"github.com/asadbek21coder/demoproject/internal/repository"
)

type Books interface {
	GetAllBooks(int, int) ([]*entities.Book, error)
}

type Authors interface {
	GetAllAuthors(int, int) ([]*entities.Author, error)
}

type Service struct {
	Books
	Authors
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Books:   NewBooksService(repos.Books),
		Authors: NewAuthorsService(repos.Authors),
	}
}
