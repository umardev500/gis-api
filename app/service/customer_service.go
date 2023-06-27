package service

import (
	"context"
	"gis/domain/model"
	"gis/domain/repository"
	"gis/domain/service"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type customerService struct {
	repo repository.CustomerRepositoryInterface
}

func NewCustomerService(repo repository.CustomerRepositoryInterface) service.CustomerServiceInterface {
	return &customerService{
		repo: repo,
	}
}

func (c *customerService) Create(ctx context.Context, payload model.CustomerRequestPayload) error {
	newPayload := payload
	newPayload.CreatedAt = time.Now().Unix()
	err := c.repo.Create(ctx, newPayload)

	return err
}

func (c *customerService) FindAll(ctx context.Context, findMeta model.FindMetaRequest) ([]model.CustomerModel, model.Meta, error) {
	filter := bson.M{}
	customers, meta, err := c.repo.FindAll(ctx, &findMeta, filter)

	return customers, meta, err
}

func (c *customerService) FindOne(ctx context.Context, id string) (*model.CustomerModel, error) {
	customer, err := c.repo.FindOne(ctx, id)

	return customer, err
}
