package repository

import (
	"github.com/asadbek21coder/demoproject/config"
	"github.com/asadbek21coder/demoproject/internal/entities"
	"github.com/asadbek21coder/demoproject/internal/logger"
	"github.com/jmoiron/sqlx"
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

type Repository struct {
	Books
	Authors
}

func NewRepository(db *sqlx.DB, log *logger.Logger, cfg *config.Config) *Repository {
	return &Repository{
		Books:   NewBooksPostgres(db, log, cfg),
		Authors: NewAuthorsPostgres(db, log, cfg),
	}
}
