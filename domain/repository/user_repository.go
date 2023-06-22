package repository

import (
	"context"
	"gis/domain/model"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepositoryInterface interface {
	FindOne(ctx context.Context, filter bson.M) (*model.UserModel, error)
}
