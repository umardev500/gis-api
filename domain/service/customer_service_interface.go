package service

import (
	"context"
	"gis/domain/model"
)

type CustomerServiceInterface interface {
	Create(ctx context.Context, payload model.CustomerRequestPayload) error
	FindOne(ctx context.Context, id string) (*model.CustomerModel, error)
	FindAll(ctx context.Context, findMeta model.FindMetaRequest) ([]model.CustomerModel, model.Meta, error)
}
