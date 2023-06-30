package repository

import (
	"context"
	"gis/domain/model"

	"go.mongodb.org/mongo-driver/bson"
)

type CustomerRepositoryInterface interface {
	Update(ctx context.Context, payload bson.M) error
	Create(ctx context.Context, payload model.CustomerRequestPayload) error
	FindOne(ctx context.Context, id string) (*model.CustomerModel, error)
	FindAll(ctx context.Context, findMeta *model.FindMetaRequest, filter bson.M) ([]model.CustomerModel, model.Meta, error)
}
