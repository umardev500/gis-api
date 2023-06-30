package service

import (
	"context"
	"gis/domain/model"

	"go.mongodb.org/mongo-driver/bson"
)

type CustomerServiceInterface interface {
	Delete(ctx context.Context, customerId string) error
	Update(ctx context.Context, payload bson.M) error
	Create(ctx context.Context, payload model.CustomerRequestPayload) error
	FindOne(ctx context.Context, id string) (*model.CustomerModel, error)
	FindAll(ctx context.Context, findMeta model.FindMetaRequest) ([]model.CustomerModel, model.Meta, error)
	FindAllNearest(ctx context.Context, findMeta model.FindMetaRequest) ([]model.CustomerModel, model.Meta, error)
}
