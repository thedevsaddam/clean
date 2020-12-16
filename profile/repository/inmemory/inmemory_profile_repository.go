package repository

import (
	"context"

	"github.com/thedevsaddam/clean/domain"
)

var inMemoryProfiles = make([]*domain.Profile, 0)

type inMemoryProfileRepository struct{}

// NewInMemoryProfileRepository ...
func NewInMemoryProfileRepository() domain.ProfileRepository {
	return new(inMemoryProfileRepository)
}

//Store ...
func (i *inMemoryProfileRepository) Store(ctx context.Context, profile *domain.Profile) (uint, error) {
	profile.ID = uint(len(inMemoryProfiles)) + 1
	inMemoryProfiles = append(inMemoryProfiles, profile)
	return profile.ID, nil
}

//GetByID ...
func (i *inMemoryProfileRepository) GetByUserID(ctx context.Context, id uint) (*domain.Profile, error) {
	if int(id) > len(inMemoryProfiles) {
		return nil, nil
	}
	for _, p := range inMemoryProfiles {
		if p.UserID == id {
			return p, nil
		}
	}
	return nil, nil
}

//Delete ...
func (i *inMemoryProfileRepository) Delete(ctx context.Context, id uint) error {
	for k, u := range inMemoryProfiles {
		if id == u.ID {
			inMemoryProfiles = append(inMemoryProfiles[:k], inMemoryProfiles[k+1:]...)
		}
	}
	return nil
}
