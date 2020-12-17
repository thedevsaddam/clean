package domain

import "context"

// Follower represents github followers
type Follower struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
}

// FollowerRepository represents the follower repository contract
type FollowerRepository interface {
	GetByUserUsername(ctx context.Context, username string) ([]*Follower, error)
}
