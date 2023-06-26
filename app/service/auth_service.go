package service

import (
	"context"
	"fmt"
	"gis/domain/model"
	"gis/domain/repository"
	"gis/domain/service"

	"go.mongodb.org/mongo-driver/bson"
)

type authService struct {
	userRepo repository.UserRepositoryInterface
}

func NewAuthService(userRepo repository.UserRepositoryInterface) service.AuthServiceInterface {
	return &authService{
		userRepo: userRepo,
	}
}

func (a *authService) Login(ctx context.Context, authCreds model.AuthCreds) (*model.UserModel, error) {
	filter := bson.M{"username": authCreds.Username, "password": authCreds.Password}
	user, err := a.userRepo.FindOne(ctx, filter)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return user, nil
}
