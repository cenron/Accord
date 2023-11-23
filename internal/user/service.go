package user

import (
	"accord/pkg/util/crypter"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	SecretKey = "change_it"
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

func (us *UserService) Signup(c context.Context, req *SignupUserReq) (*SignupUserRes, error) {
	ctx, cancel := context.WithTimeout(c, us.timeout)
	defer cancel()

	// Generate our hashed password.
	hashPassword, err := crypter.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	var u = &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashPassword,
	}

	user, err := us.store.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &SignupUserRes{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	return res, nil
}

type JWTClaim struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (us *UserService) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, us.timeout)
	defer cancel()

	user, err := us.store.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if err := crypter.CheckPassword(req.Password, user.Password); err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaim{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.ID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return nil, err
	}

	res := &LoginUserRes{
		accessToken: ss,
		ID:          user.ID,
		Username:    user.Username,
	}

	return res, nil
}
