package repository

import (
	"github.com/asadbek21coder/demoproject/internal/entities"
	"github.com/jmoiron/sqlx"
)

type Books interface {
	GetAllBooks() ([]entities.Book, error)
}

type Repository struct {
	Books
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Books: NewBooksPostgres(db),
	}
}
