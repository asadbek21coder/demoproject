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

func (s *AuthorsService) CreateAuthor(req *entities.Author) (*entities.Author, error) {
	return s.repo.CreateAuthor(req)
}

func (s *AuthorsService) GetAuthorById(authorId int) (*entities.Author, error) {
	return s.repo.GetAuthorById(authorId)
}

func (s *AuthorsService) UpdateAuthor(req *entities.Author) (*entities.Author, error) {
	return s.repo.UpdateAuthor(req)
}
