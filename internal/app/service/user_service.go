package service

import (
	"context"
	"errors"

	"github.com/sk-pathak/go-structure/internal/app/repository"
	"github.com/sk-pathak/go-structure/internal/db"
)

type UserService struct {
	repo *repository.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx context.Context, user *db.User) error {
	if err := s.repo.SaveUser(ctx, user); err != nil {
		return errors.New("failed to create user in repository: " + err.Error())
	}
	return nil
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers(ctx context.Context) ([]db.User, error) {
	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, errors.New("failed to retrieve users from repository: " + err.Error())
	}
	return users, nil
}
