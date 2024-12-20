package repository

import (
	"context"

	"github.com/sk-pathak/go-structure/internal/db"
)

// UserRepository interacts with the database for user-related operations.
type UserRepository struct {
	queries *db.Queries
}

// NewUserRepository creates a new UserRepository using the generated queries.
func NewUserRepository(queries *db.Queries) *UserRepository {
	return &UserRepository{queries: queries}
}

// SaveUser saves a user to the database.
func (r *UserRepository) SaveUser(ctx context.Context, user *db.User) error {
	// Call the generated method to insert the user into the database
	_, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Name:  user.Name,
		Email: user.Email,
		// Populate other fields as required by your CreateUserParams
	})
	return err
}

// GetAllUsers retrieves all users from the database.
func (r *UserRepository) GetAllUsers(ctx context.Context) ([]db.User, error) {
	// Call the generated method to get all users from the database
	users, err := r.queries.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the result from db.Queries (which might be db.User) to model.User if needed
	var result []db.User
	for _, u := range users {
		result = append(result, db.User{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
			// Map other fields as needed
		})
	}

	return result, nil
}
