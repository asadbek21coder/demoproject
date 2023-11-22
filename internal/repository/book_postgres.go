package repository

import (
	"fmt"

	"github.com/asadbek21coder/demoproject/internal/entities"
	"github.com/jmoiron/sqlx"
)

const (
	bookTable = "books"
)

type BooksPostgres struct {
	db *sqlx.DB
}

func NewBooksPostgres(db *sqlx.DB) *BooksPostgres {
	return &BooksPostgres{db: db}
}

func (r *BooksPostgres) GetAllBooks() ([]entities.Book, error) {
	var result []entities.Book
	query := fmt.Sprintf("SELECT id, name, author, price FROM %s ", bookTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return []entities.Book{}, err
	}

	for rows.Next() {
		var book entities.Book
		err := rows.Scan(
			book.ID,
			book.Name,
			book.Author,
			book.Price,
		)
		if err != nil {
			return []entities.Book{}, err
		}

		result = append(result, book)
	}

	return result, err
}
