package infrastructure

import (
	"errors"
	"sync"

	"github.com/meseger/portifolio/wire/internal/user/domain"
)

type MockUserRepository struct {
	users map[string]*domain.User
	mu    sync.RWMutex
	// You can add fields to track method calls for testing
	SaveCalled     bool
	FindByIDCalled bool
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *MockUserRepository) Save(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.SaveCalled = true
	r.users[user.ID] = user
	return nil
}

func (r *MockUserRepository) FindByID(id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	r.FindByIDCalled = true
	if user, exists := r.users[id]; exists {
		return user, nil
	}
	return nil, errors.New("user not found")
}
