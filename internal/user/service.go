package user

import (
	"context"
	"time"
)

type UserService struct {
	store   *UserRepository
	timeout time.Duration
}

func NewUserService(store UserStore) *UserService {
	return &UserService{
		NewUserRepository(store),
		time.Duration(2) * time.Second,
	}
}

func (us *UserService) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, us.timeout)
	defer cancel()

	var u = &User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := us.store.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	return res, nil
}