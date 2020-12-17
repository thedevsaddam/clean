package presenter

import "github.com/thedevsaddam/clean/domain"

//Presenter ...
type Presenter interface {
	PresentUser(*domain.User) *User
	PresentUsers([]*domain.User) []*User
}

//NewUserPresenter ...
func NewUserPresenter() *UserPresenter {
	return new(UserPresenter)
}

// UserPresenter ...
type UserPresenter struct{}

// User ...
type User struct {
	ID        uint               `json:"id"`
	Username  string             `json:"username"`
	Type      string             `json:"type"`
	Profile   *domain.Profile    `json:"profile,omitempty"`
	Followers []*domain.Follower `json:"followers,omitempty"`
}

//PresentUser ...
func (*UserPresenter) PresentUser(u *domain.User) *User {
	return &User{
		ID:        u.ID,
		Username:  u.Username,
		Type:      u.Type,
		Profile:   u.Profile,
		Followers: u.Followers,
	}
}

//PresentUsers ...
func (u *UserPresenter) PresentUsers(usr []*domain.User) []*User {
	var users []*User
	for _, uv := range usr {
		users = append(users, u.PresentUser(uv))
	}
	if len(users) == 0 {
		return []*User{} // if no data send empty; not nil
	}
	return users
}
