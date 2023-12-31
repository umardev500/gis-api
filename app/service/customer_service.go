package service

import (
	"context"
	"gis/domain/model"
	"gis/domain/repository"
	"gis/domain/service"
	"strconv"
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

func (c *customerService) Delete(ctx context.Context, customerId string) error {
	err := c.repo.Delete(ctx, customerId)
	return err
}

func (c *customerService) Update(ctx context.Context, payload bson.M, customerID string) error {
	err := c.repo.Update(ctx, payload, customerID)

	return err
}

func (c *customerService) Create(ctx context.Context, payload model.CustomerRequestPayload) error {
	newPayload := payload
	now := time.Now()
	newPayload.ID = strconv.Itoa(int(now.UnixMilli()))
	newPayload.CreatedAt = now.Unix()
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

func (c *customerService) FindAllNearest(ctx context.Context, findMeta model.FindMetaRequest) ([]model.CustomerModel, model.Meta, error) {
	filter := bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{105.87665435422802, -6.5186988908316055},
				},
			},
		},
	}

	customers, meta, err := c.repo.FindAll(ctx, &findMeta, filter)

	return customers, meta, err
}
