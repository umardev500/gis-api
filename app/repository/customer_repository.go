package repository

import (
	"context"
	"fmt"
	"gis/domain/model"
	"gis/domain/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type customerRepository struct {
	customerCollection *mongo.Collection
}

func NewCustomerRepository(db *mongo.Database) repository.CustomerRepositoryInterface {
	return &customerRepository{
		customerCollection: db.Collection("customers"),
	}
}

func (c *customerRepository) Create(ctx context.Context, payload model.CustomerRequestPayload) error {
	_, err := c.customerCollection.InsertOne(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) FindAll(ctx context.Context) ([]model.CustomerModel, model.Meta, error) {
	filter := bson.M{}
	cur, err := c.customerCollection.Find(ctx, filter)
	if err != nil {
		return nil, model.Meta{}, err
	}

	defer cur.Close(ctx)

	var customers []model.CustomerModel

	for cur.Next(ctx) {
		var each model.CustomerModel
		if err := cur.Decode(&each); err != nil {
			fmt.Println("Error decoding customer:", err)
			continue
		}

		customers = append(customers, each)
	}

	total, _ := c.customerCollection.CountDocuments(ctx, filter)

	meta := model.Meta{
		Total: total,
	}

	return customers, meta, nil
}

func (c *customerRepository) FindOne(ctx context.Context, id string) (*model.CustomerModel, error) {
	var customer model.CustomerModel
	filter := bson.M{"id": id}
	if err := c.customerCollection.FindOne(ctx, filter).Decode(&customer); err != nil {
		return nil, err
	}

	return &customer, nil
}
