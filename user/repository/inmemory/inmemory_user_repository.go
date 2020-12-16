package repository

import (
	"context"

	"github.com/thedevsaddam/clean/domain"
)

var inUsers = make([]*domain.User, 0)

// inMemoryUserRepository ...
type inMemoryUserRepository struct {
}

// NewInMemoryUserRepository ...
func NewInMemoryUserRepository() domain.UserRepository {
	return new(inMemoryUserRepository)
}

// Store store user in in-memory
func (i *inMemoryUserRepository) Store(ctx context.Context, user *domain.User) (uint, error) {
	user.ID = uint(len(inUsers)) + 1
	inUsers = append(inUsers, user)
	return user.ID, nil
}

// Fetch return users from in-memory
func (i *inMemoryUserRepository) Fetch(ctx context.Context, ctr *domain.UserCriteria) ([]*domain.User, error) {
	uu := make([]*domain.User, 0)
	// skip filtering
	for _, u := range inUsers {
		uu = append(uu, u)
	}
	return uu, nil
}

// GetByID return user from in-memory by its matching id
func (i *inMemoryUserRepository) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	if int(id) > len(inUsers) {
		return nil, nil
	}
	for _, u := range inUsers {
		if id == u.ID {
			return u, nil
		}
	}
	return nil, nil
}

// Delete remove a user from in-memory by its matching id
func (i *inMemoryUserRepository) Delete(ctx context.Context, id uint) error {
	for k, u := range inUsers {
		if id == u.ID {
			inUsers = append(inUsers[:k], inUsers[k+1:]...)
		}
	}
	return nil
}
