package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

const CollectionName = "users"

type UserStore interface {
	InsertOne(ctx context.Context, collection string, data []byte) (string, error)
	GetById(ctx context.Context, collection string, id string) ([]byte, error)
}

type UserRepository struct {
	store UserStore
}

func NewUserRepository(store UserStore) *UserRepository {
	return &UserRepository{
		store: store,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *User) (*User, error) {
	// Try to convert our object into json.
	userJson, err := bson.Marshal(user)
	if err != nil {
		return nil, err
	}

	lastId, err := r.store.InsertOne(ctx, CollectionName, userJson)
	if err != nil {
		return &User{}, err
	}

	user.ID = lastId
	return user, nil
}

func (r *UserRepository) GetUserById(ctx context.Context, id string) (*User, error) {
	// Get the ID from the collection.
	result, err := r.store.GetById(ctx, CollectionName, id)
	if err != nil {
		return &User{}, err
	}

	var user User
	if err := bson.Unmarshal(result, &user); err != nil {
		return &User{}, err
	}

	return &user, nil
}
