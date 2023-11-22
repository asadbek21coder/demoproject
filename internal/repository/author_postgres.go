package repository

import (
	"database/sql"
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
	authorTable = "author"
)

type AuthorsPostgres struct {
	Log *logger.Logger
	Cfg *config.Config
	db  *sqlx.DB
}

func NewAuthorsPostgres(db *sqlx.DB, log *logger.Logger, cfg *config.Config) *AuthorsPostgres {
	return &AuthorsPostgres{
		db:  db,
		Log: log,
		Cfg: cfg,
	}
}

func (a *AuthorsPostgres) GetAllAuthors(page, limit int) ([]*entities.Author, error) {
	offset := (page - 1) * limit
	result := []*entities.Author{}
	query := fmt.Sprintf(`

	SELECT 
		author_id, 
		author_name, 
		created_at,
		updated_at
	FROM  
		%s
	WHERE 
		deleted_at IS NULL
	LIMIT $1 OFFSET $2
	`, authorTable)

	rows, err := a.db.Query(query, limit, offset)
	if err != nil {
		a.Log.Error(err.Error())
		return []*entities.Author{}, status.Error(codes.Internal, "Ooops something went wrong")
	}

	for rows.Next() {
		author := &entities.Author{}
		err := rows.Scan(
			&author.ID,
			&author.Name,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return []*entities.Author{}, status.Error(codes.Internal, "Oops something went wrong.")
		}
		author.CreatedAt = CreatedAt.Format(time.RFC3339)
		author.UpdatedAt = UpdatedAt.Format(time.RFC3339)

		result = append(result, author)
	}

	return result, err
}

func (a *AuthorsPostgres) GetAuthorById(id int) (*entities.Author, error) {
	author := &entities.Author{}
	query := `
		SELECT 
			author_id,
			author_name, 
			created_at, 
			updated_at 
		FROM 
			author 
		WHERE 
			author_id=$1 
		AND 
			deleted_at 
		IS NULL`

	err := a.db.QueryRow(query, id).Scan(
		&author.ID,
		&author.Name,
		&CreatedAt,
		&UpdatedAt,
	)

	author.CreatedAt = CreatedAt.Format(time.RFC1123)
	author.UpdatedAt = UpdatedAt.Format(time.RFC1123)

	if err == sql.ErrNoRows {
		a.Log.Error(err.Error())
		return &entities.Author{}, status.Error(codes.NotFound, "This author doesn't exist.")
	}
	if err != nil {
		a.Log.Error(err.Error())
		return &entities.Author{}, status.Error(codes.Internal, "Oops something went wrong.")
	}

	return author, nil
}

func (a *AuthorsPostgres) CreateAuthor(req *entities.CreateAuthorReq) (*entities.Author, error) {
	var res = &entities.Author{}
	query := `
	INSERT INTO author (
		author_name
		) 
	VALUES
		($1) 
	RETURNING 
		author_id,
		author_name, 
		created_at, 
		updated_at;
		`

	err := a.db.QueryRow(
		query,
		req.Name).Scan(
		&res.ID,
		&res.Name,
		&CreatedAt,
		&UpdatedAt)

	if err != nil {
		a.Log.Error(err.Error())
		return &entities.Author{}, status.Error(codes.Internal, "Ooops something went wrong")
	}
	res.CreatedAt = CreatedAt.Format(time.RFC1123)
	res.UpdatedAt = UpdatedAt.Format(time.RFC1123)
	return res, nil
}

func (a *AuthorsPostgres) UpdateAuthor(id int, req *entities.UpdateAuthorReq) (*entities.Author, error) {
	var res = &entities.Author{}
	query := `
	UPDATE 
		author 
	SET 
		author_name =$1,
		updated_at = now()
	WHERE 
		author_id=$2 
	AND 
		deleted_at 
	IS NULL 
	RETURNING 
		author_id, 
		author_name, 
		created_at,
		updated_at;`

	err := a.db.QueryRow(
		query, req.Name, id).Scan(
		&res.ID,
		&res.Name,
		&CreatedAt,
		&UpdatedAt)

	if err != nil {
		a.Log.Error(err.Error())
		return &entities.Author{}, status.Error(codes.Internal, "Ooops something went wrong")
	}

	res.CreatedAt = CreatedAt.Format(time.RFC1123)
	res.UpdatedAt = UpdatedAt.Format(time.RFC1123)

	return res, nil
}

func (a *AuthorsPostgres) DeleteAuthor(authorId int) error {
	queryDeleteAuthor := `
	UPDATE 
		author 
	SET 
		deleted_at = now() 
	WHERE 
		author_2id=$1`

	_, err := a.db.Exec(queryDeleteAuthor, authorId)
	if err != nil {
		a.Log.Error(err.Error())
		return status.Error(codes.Internal, "Ooops something went wrong")
	}
	return nil
}
