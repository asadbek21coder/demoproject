package service

import (
	"github.com/asadbek21coder/demoproject/internal/entities"
	"github.com/asadbek21coder/demoproject/internal/repository"
)

type AuthorsService struct {
	repo repository.Authors
}

func NewAuthorsService(repo repository.Authors) *AuthorsService {
	return &AuthorsService{repo: repo}

}

func (s *AuthorsService) GetAllAuthors(page, limit int) ([]*entities.Author, error) {
	return s.repo.GetAllAuthors(page, limit)
}
