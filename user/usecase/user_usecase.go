package usecase

import (
	"context"

	"github.com/thedevsaddam/clean/domain"
)

// userUsecase ...
type userUsecase struct {
	userRepository     domain.UserRepository
	profileRepository  domain.ProfileRepository
	followerRepository domain.FollowerRepository
}

// NewUserUsecase ...
func NewUserUsecase(u domain.UserRepository, p domain.ProfileRepository, f domain.FollowerRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository:     u,
		profileRepository:  p,
		followerRepository: f,
	}
}

// Store ...
func (u *userUsecase) Store(ctx context.Context, user *domain.User) (uint, error) {
	uid, err := u.userRepository.Store(ctx, user)
	if err != nil {
		return 0, err
	}

	user.Profile.UserID = uid
	_, err = u.profileRepository.Store(ctx, user.Profile)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

// Fetch ...
func (u *userUsecase) Fetch(ctx context.Context, ctr *domain.UserCriteria) ([]*domain.User, error) {
	users, err := u.userRepository.Fetch(ctx, ctr)
	if err != nil {
		return nil, err
	}

	for k, v := range users {
		p, _ := u.profileRepository.GetByUserID(ctx, v.ID)
		users[k].Profile = p
		f, _ := u.followerRepository.GetByUserUsername(ctx, v.Username)
		users[k].Followers = f
	}
	return users, nil
}

// GetByID ..
func (u *userUsecase) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	user, err := u.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	profile, err := u.profileRepository.GetByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	user.Profile = profile

	f, err := u.followerRepository.GetByUserUsername(ctx, user.Username)
	if err != nil {
		return nil, err
	}
	user.Followers = f

	return user, nil
}

// Delete ....
func (u *userUsecase) Delete(ctx context.Context, id uint) error {
	if err := u.profileRepository.Delete(ctx, id); err != nil {
		return err
	}
	return u.Delete(ctx, id)
}
