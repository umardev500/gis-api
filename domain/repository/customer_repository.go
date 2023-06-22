package repository

import (
	"context"
	"gis/domain/model"
)

type CustomerRepositoryInterface interface {
	Create(ctx context.Context, payload model.CustomerRequestPayload) error
	FindAll(ctx context.Context) ([]model.CustomerModel, model.Meta, error)
}
