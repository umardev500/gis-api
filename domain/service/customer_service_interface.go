package service

import (
	"context"
	"gis/domain/model"
)

type CustomerServiceInterface interface {
	Create(ctx context.Context, payload model.CustomerRequestPayload) error
	FindAll(ctx context.Context) ([]model.CustomerModel, model.Meta, error)
}
