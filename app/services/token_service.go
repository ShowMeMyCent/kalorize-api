package services

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"kalorize-api/app/models"
	"kalorize-api/app/repositories"
)

type tokenService struct {
	repo repositories.TokenRepository
}

func NewTokenService(db *gorm.DB) *tokenService {
	return &tokenService{repo: repositories.NewDBTokenRepository(db)}
}

func (s *tokenService) GetAllTokens() ([]models.Token, error) {
	tokens, err := s.repo.GetToken()
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (s *tokenService) GetTokenByUserEmail(IdUser string, idToken string) (*models.Token, error) {
	token, err := s.repo.GetTokenByUserEmail(IdUser, idToken)
	if err != nil {
		return nil, err
	}
	if token == nil {
		return nil, errors.New("token not found")
	}
	return token, nil
}

func (s *tokenService) CreateToken(token models.Token) error {
	return s.repo.CreateNewToken(token)
}

func (s *tokenService) UpdateToken(token models.Token) error {
	return s.repo.UpdateToken(token)
}

func (s *tokenService) DeleteToken(idToken string) error {
	_, err := uuid.Parse(idToken)
	if err != nil {
		return err
	}
	return s.repo.DeleteToken(idToken)
}

type TokenService interface {
	GetAllTokens() ([]models.Token, error)
	GetTokenByUserEmail(email string, idToken string) (*models.Token, error)
	CreateToken(token models.Token) error
	UpdateToken(token models.Token) error
	DeleteToken(idToken string) error
}
