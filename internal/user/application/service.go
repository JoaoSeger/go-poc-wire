package application

import (
	"github.com/google/uuid"
	"github.com/meseger/portifolio/wire/internal/user/domain"
)

type UserService struct {
	repository domain.UserRepository
}

func NewUserService(repository domain.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) CreateUser(name, email string) (*domain.User, error) {
	user := &domain.User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}

	if err := s.repository.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUser(id string) (*domain.User, error) {
	return s.repository.FindByID(id)
}
