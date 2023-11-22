package repository

import (
	"fmt"
	"time"

	"github.com/asadbek21coder/demoproject/config"
	"github.com/asadbek21coder/demoproject/internal/entities"
	"github.com/asadbek21coder/demoproject/internal/logger"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	bookTable = "book"
)

var (
	CreatedAt time.Time
	UpdatedAt time.Time
)

type BooksPostgres struct {
	Log *logger.Logger
	Cfg *config.Config
	db  *sqlx.DB
}

func NewBooksPostgres(db *sqlx.DB, log *logger.Logger, cfg *config.Config) *BooksPostgres {
	return &BooksPostgres{
		db:  db,
		Log: log,
		Cfg: cfg,
	}
}

func (r *BooksPostgres) GetAllBooks(page, limit int) ([]*entities.Book, error) {
	offset := (page - 1) * limit
	result := []*entities.Book{}
	query := fmt.Sprintf(`

	SELECT 
		book_id, 
		book_title, 
		book_author, 
		book_price,
		created_at,
		updated_at
	FROM  
		%s
	WHERE 
		deleted_at IS NULL
	LIMIT $1 OFFSET $2
	`, bookTable)

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		r.Log.Error(err.Error())
		return []*entities.Book{}, status.Error(codes.Internal, "Ooops something went wrong")
	}

	for rows.Next() {
		book := &entities.Book{}
		err := rows.Scan(
			&book.ID,
			&book.Name,
			&book.Author,
			&book.Price,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return []*entities.Book{}, status.Error(codes.Internal, "Oops something went wrong.")
		}
		book.CreatedAt = CreatedAt.Format(time.RFC3339)
		book.UpdatedAt = UpdatedAt.Format(time.RFC3339)

		result = append(result, book)
	}

	return result, err
}
