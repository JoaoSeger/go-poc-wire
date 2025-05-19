package application

import (
	"testing"

	"github.com/meseger/portifolio/wire/internal/user/infrastructure"
)

func TestUserService_CreateUser(t *testing.T) {
	// Arrange
	repo := infrastructure.NewMockUserRepository()
	service := NewUserService(repo)

	// Act
	user, err := service.CreateUser("John Doe", "john@example.com")

	// Assert
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if user.Name != "John Doe" {
		t.Errorf("expected name to be John Doe, got %s", user.Name)
	}
	if user.Email != "john@example.com" {
		t.Errorf("expected email to be john@example.com, got %s", user.Email)
	}
}

func TestUserService_GetUser(t *testing.T) {
	// Arrange
	repo := infrastructure.NewMockUserRepository()
	service := NewUserService(repo)

	// Create a user first
	createdUser, err := service.CreateUser("John Doe", "john@example.com")
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// Act
	foundUser, err := service.GetUser(createdUser.ID)

	// Assert
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if foundUser.ID != createdUser.ID {
		t.Errorf("expected ID to be %s, got %s", createdUser.ID, foundUser.ID)
	}
}
