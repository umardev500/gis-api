package service

import (
	"context"
	"gis/domain/model"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, authCreds model.AuthCreds) (*model.UserModel, error)
}
