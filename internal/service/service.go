package service

import (
	"github.com/asadbek21coder/demoproject/internal/entities"
	"github.com/asadbek21coder/demoproject/internal/repository"
)

type Books interface {
	GetAllBooks(int, int) ([]*entities.Book, error)
	GetBookById(int) (*entities.Book, error)
	UpdateBook(int, *entities.UpdateBookReq) (*entities.Book, error)
	CreateBook(*entities.CreateBookReq) (*entities.Book, error)
	DeleteBook(int) error
}

type Authors interface {
	GetAllAuthors(int, int) ([]*entities.Author, error)
	GetAuthorById(int) (*entities.Author, error)
	UpdateAuthor(int, *entities.UpdateAuthorReq) (*entities.Author, error)
	CreateAuthor(*entities.CreateAuthorReq) (*entities.Author, error)
	DeleteAuthor(int) error
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
