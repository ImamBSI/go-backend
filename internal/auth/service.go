// Business logic Autentikasi: register, login, token generation, dll.
package auth

import (
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(username, password string, accountID uint) error {
	existing, err := s.repo.FindByUsername(username)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.New("username already exists")
	}

	hashed, err := HashPassword(password)
	if err != nil {
		return err
	}

	user := &User{
		Username:  username,
		Password:  hashed,
		AccountID: accountID,
	}

	return s.repo.Create(user)
}

func (s *Service) Login(username, password string) (string, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("invalid credentials")
	}

	if err := ComparePassword(user.Password, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
