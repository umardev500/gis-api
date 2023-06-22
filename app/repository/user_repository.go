package repository

import (
	"context"
	"gis/domain/model"
	"gis/domain/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	customerCollection *mongo.Collection
}

func NewUserRepository() repository.UserRepositoryInterface {
	return &userRepository{}
}

func (u *userRepository) FindOne(ctx context.Context, filter bson.M) (*model.UserModel, error) {
	var user model.UserModel

	if err := u.customerCollection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
