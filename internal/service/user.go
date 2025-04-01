package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/vasujain275/bookbridge-api/internal/repository"
)

// UserServiceImpl implements the UserService interface
type UserServiceImpl struct {
	repo *repository.Queries
}

// NewUserService creates a new user service
func NewUserService(repo *repository.Queries) UserService {
	return &UserServiceImpl{
		repo: repo,
	}
}

// GetByID gets a user by ID
func (s *UserServiceImpl) GetByID(ctx context.Context, id uuid.UUID) (*repository.User, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

// GetByUsername gets a user by username
func (s *UserServiceImpl) GetByUsername(ctx context.Context, username string) (*repository.User, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}
	return &user, nil
}

// GetByEmail gets a user by email
func (s *UserServiceImpl) GetByEmail(ctx context.Context, email string) (*repository.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return &user, nil
}

// List gets a list of users
func (s *UserServiceImpl) List(ctx context.Context, limit, offset int32) ([]*repository.User, error) {
	users, err := s.repo.ListUsers(ctx, repository.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	// Convert []repository.User to []*repository.User
	userPtrs := make([]*repository.User, len(users))
	for i := range users {
		userPtrs[i] = &users[i]
	}

	return userPtrs, nil
}

// Create creates a new user
func (s *UserServiceImpl) Create(ctx context.Context, params repository.CreateUserParams) (*repository.User, error) {
	// Check if user with username already exists
	_, err := s.repo.GetUserByUsername(ctx, params.Username)
	if err == nil {
		return nil, errors.New("username already exists")
	}

	// Check if user with email already exists
	_, err = s.repo.GetUserByEmail(ctx, params.Email)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	user, err := s.repo.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &user, nil
}

// Update updates a user
func (s *UserServiceImpl) Update(ctx context.Context, params repository.UpdateUserParams) (*repository.User, error) {
	// Check if user exists
	_, err := s.repo.GetUser(ctx, params.ID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	user, err := s.repo.UpdateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	return &user, nil
}

// Delete deletes a user
func (s *UserServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
