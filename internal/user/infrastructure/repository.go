package infrastructure

import (
	"errors"
	"sync"

	"github.com/meseger/portifolio/wire/internal/user/domain"
)

type InMemoryUserRepository struct {
	users map[string]*domain.User
	mu    sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *InMemoryUserRepository) Save(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) FindByID(id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if user, exists := r.users[id]; exists {
		return user, nil
	}
	return nil, errors.New("user not found")
}
