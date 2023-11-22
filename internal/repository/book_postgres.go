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

func (a *BooksPostgres) GetBookById(id int) (*entities.Book, error) {
	book := &entities.Book{}
	// query := `
	// 	SELECT
	// 		author_id,
	// 		author_name,
	// 		created_at,
	// 		updated_at
	// 	FROM
	// 		author
	// 	WHERE
	// 		author_id=$1
	// 	AND
	// 		deleted_at
	// 	IS NULL`

	// err := a.db.QueryRow(query, id).Scan(
	// 	&author.ID,
	// 	&author.Name,
	// 	&CreatedAt,
	// 	&UpdatedAt,
	// )

	// author.CreatedAt = CreatedAt.Format(time.RFC1123)
	// author.UpdatedAt = UpdatedAt.Format(time.RFC1123)

	// if err == sql.ErrNoRows {
	// 	a.Log.Error(err.Error())
	// 	return &entities.Book{}, status.Error(codes.NotFound, "This author doesn't exist.")
	// }
	// if err != nil {
	// 	a.Log.Error(err.Error())
	// 	return &entities.Book{}, status.Error(codes.Internal, "Oops something went wrong.")
	// }

	return book, nil
}

func (a *BooksPostgres) CreateBook(req *entities.CreateBookReq) (*entities.Book, error) {
	var res = &entities.Book{}
	// query := `
	// INSERT INTO author (
	// 	author_name
	// 	)
	// VALUES
	// 	($1)
	// RETURNING
	// 	author_id,
	// 	author_name,
	// 	created_at,
	// 	updated_at;
	// 	`

	// err := a.db.QueryRow(
	// 	query,
	// 	req.Name).Scan(
	// 	&res.ID,
	// 	&res.Name,
	// 	&CreatedAt,
	// 	&UpdatedAt)

	// if err != nil {
	// 	a.Log.Error(err.Error())
	// 	return &entities.Book{}, status.Error(codes.Internal, "Ooops something went wrong")
	// }
	// res.CreatedAt = CreatedAt.Format(time.RFC1123)
	// res.UpdatedAt = UpdatedAt.Format(time.RFC1123)
	return res, nil
}

func (a *BooksPostgres) UpdateBook(id int, req *entities.UpdateBookReq) (*entities.Book, error) {
	var res = &entities.Book{}
	// query := `
	// UPDATE
	// 	author
	// SET
	// 	author_name =$1,
	// 	updated_at = now()
	// WHERE
	// 	author_id=$2
	// AND
	// 	deleted_at
	// IS NULL
	// RETURNING
	// 	author_id,
	// 	author_name,
	// 	created_at,
	// 	updated_at;`

	// err := a.db.QueryRow(
	// 	query, req.Name, id).Scan(
	// 	&res.ID,
	// 	&res.Name,
	// 	&CreatedAt,
	// 	&UpdatedAt)

	// if err != nil {
	// 	a.Log.Error(err.Error())
	// 	return &entities.Book{}, status.Error(codes.Internal, "Ooops something went wrong")
	// }

	// res.CreatedAt = CreatedAt.Format(time.RFC1123)
	// res.UpdatedAt = UpdatedAt.Format(time.RFC1123)

	return res, nil
}

func (a *BooksPostgres) DeleteBook(bookId int) error {
	// queryDeleteBook := `
	// UPDATE
	// 	author
	// SET
	// 	deleted_at = now()
	// WHERE
	// 	author_2id=$1`

	// _, err := a.db.Exec(queryDeleteBook, authorId)
	// if err != nil {
	// 	a.Log.Error(err.Error())
	// 	return status.Error(codes.Internal, "Ooops something went wrong")
	// }
	return nil
}
