package domain

import (
	"context"
	"time"
)

// Profile represents an user's profile
type Profile struct {
	ID        uint       `json:"id"`
	UserID    uint       `json:"user_id"`
	Name      string     `json:"name"`
	Age       uint       `json:"age"`
	Bio       string     `json:"bio"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// ProfileRepository represent the users's repository contract
type ProfileRepository interface {
	Store(ctx context.Context, profile *Profile) (uint, error)
	GetByUserID(ctx context.Context, id uint) (*Profile, error)
	Delete(ctx context.Context, id uint) error
}
