package domain

import (
	"context"
	"time"
)

// User represents user data model
type User struct {
	ID        uint       `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Type      string     `json:"type"`
	Profile   *Profile   `json:"profile"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// UserCriteria ...
type UserCriteria struct {
	Name          *string
	Offset, Limit int
}

// UserRepository ...
type UserRepository interface {
	Store(ctx context.Context, user *User) (uint, error)
	Fetch(ctx context.Context, ctr *UserCriteria) ([]*User, error)
	GetByID(ctx context.Context, id uint) (*User, error)
	Delete(ctx context.Context, id uint) error
}

// UserUsecase ...
type UserUsecase interface {
	Store(ctx context.Context, user *User) (uint, error)
	Fetch(ctx context.Context, ctr *UserCriteria) ([]*User, error)
	GetByID(ctx context.Context, id uint) (*User, error)
	Delete(ctx context.Context, id uint) error
}
