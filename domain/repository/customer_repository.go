package repository

import (
	"context"
	"gis/domain/model"
)

type CustomerRepositoryInterface interface {
	Create(ctx context.Context, payload model.CustomerRequestPayload) error
	FindOne(ctx context.Context, id string) (*model.CustomerModel, error)
	FindAll(ctx context.Context) ([]model.CustomerModel, model.Meta, error)
}
