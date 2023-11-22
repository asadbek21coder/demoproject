package service

import (
	"github.com/asadbek21coder/demoproject/internal/entities"
	"github.com/asadbek21coder/demoproject/internal/repository"
)

type Books interface {
	GetAllBooks() ([]entities.Book, error)
}

type Service struct {
	Books
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Books: NewBooksService(repos.Books),
	}
}
