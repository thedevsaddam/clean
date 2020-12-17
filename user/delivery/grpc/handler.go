package grpc

import (
	context "context"

	"github.com/thedevsaddam/clean/domain"
	grpc "google.golang.org/grpc"
)

type userGRPC struct {
	UnimplementedUserServiceServer

	UserUsecase domain.UserUsecase
}

// NewUserGRPCHandler will initialize the resources grpc endpoint
func NewUserGRPCHandler(s *grpc.Server, us domain.UserUsecase) {
	handler := userGRPC{
		UserUsecase: us,
	}
	RegisterUserServiceServer(s, &handler)
}

func (u *userGRPC) Store(ctx context.Context, in *ReqCreateUser) (*ResCreateUser, error) {
	resp := &ResCreateUser{}
	if in.Profile == nil {
		resp.Message = "User's profile data required"
		return resp, nil
	}
	id, err := u.UserUsecase.Store(ctx, &domain.User{
		Username: in.Username,
		Password: in.Password,
		Type:     in.Type,
		Profile: &domain.Profile{
			Name: in.Profile.Name,
			Age:  uint(in.Profile.Age),
			Bio:  in.Profile.Bio,
		},
	})

	if err != nil {
		resp.Message = "Failed to create user"
		return resp, nil
	}

	resp.Message = "User created successfully"
	resp.Id = uint64(id)

	return resp, nil
}

func (u *userGRPC) FetchUsers(ctx context.Context, in *ReqFetchUsers) (*RespFetchUsers, error) {
	ctr := &domain.UserCriteria{}
	users, err := u.UserUsecase.Fetch(ctx, ctr)
	resp := &RespFetchUsers{}
	if err != nil {
		resp.Message = "Failed to fetch users"
		return resp, nil
	}
	respUsers := make([]*User, 0)
	for _, u := range users {
		usr := &User{
			Id:       uint64(u.ID),
			Username: u.Username,
			Type:     u.Type,
		}

		if u.Profile != nil {
			usr.Profile = &Profile{
				Name: u.Profile.Name,
				Age:  int32(u.Profile.Age),
				Bio:  u.Profile.Bio,
			}
		}

		if u.Followers != nil {
			for _, f := range u.Followers {
				usr.Followers = append(usr.Followers, &Follower{
					Id:    int32(f.ID),
					Login: f.Login,
				})
			}
		}

		respUsers = append(respUsers, usr)
	}

	resp.Users = respUsers

	return resp, nil
}

func (u *userGRPC) FetchUser(ctx context.Context, in *ReqFetchUser) (*RespFetchUser, error) {
	user, err := u.UserUsecase.GetByID(ctx, uint(in.Id))
	resp := &RespFetchUser{}
	if err != nil {
		resp.Message = "Failed to fetch user"
		return resp, nil
	}

	if user == nil {
		resp.Message = "User not found"
		return resp, nil
	}

	resp.User = &User{
		Id:       uint64(user.ID),
		Username: user.Username,
	}

	if user.Profile != nil {
		resp.User.Profile = &Profile{
			Name: user.Profile.Name,
			Age:  int32(user.Profile.Age),
			Bio:  user.Profile.Bio,
		}
	}

	if user.Followers != nil {
		for _, f := range user.Followers {
			resp.User.Followers = append(resp.User.Followers, &Follower{
				Id:    int32(f.ID),
				Login: f.Login,
			})
		}
	}

	return resp, nil
}
