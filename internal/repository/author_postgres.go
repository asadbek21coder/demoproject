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

func (r *AuthorsPostgres) GetAllAuthors(page, limit int) ([]*entities.Author, error) {
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

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		r.Log.Error(err.Error())
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
