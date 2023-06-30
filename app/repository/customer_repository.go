package repository

import (
	"context"
	"fmt"
	"gis/domain/model"
	"gis/domain/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type customerRepository struct {
	customerCollection *mongo.Collection
}

func NewCustomerRepository(db *mongo.Database) repository.CustomerRepositoryInterface {
	return &customerRepository{
		customerCollection: db.Collection("customers"),
	}
}

func (c *customerRepository) Update(ctx context.Context, payload bson.M) error {
	filter := bson.M{"id": 1}
	update := bson.M{"$set": payload}

	_, err := c.customerCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	// fmt.Println(updateResult)

	return nil
}

func (c *customerRepository) Create(ctx context.Context, payload model.CustomerRequestPayload) error {
	_, err := c.customerCollection.InsertOne(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) FindAll(ctx context.Context, findMeta *model.FindMetaRequest, filter bson.M) ([]model.CustomerModel, model.Meta, error) {
	findOptions := options.Find()
	if findMeta != nil {
		if findMeta.PerPage != 0 {
			findOptions.SetLimit(findMeta.PerPage)
		}
		if findMeta.Order != "" && (findMeta.Order == "desc" || findMeta.Order == "asc") {
			if findMeta.Order == "desc" {
				fmt.Println("order:", findMeta.Order)
				findOptions.SetSort(bson.M{"createdAt": -1})
			}
		}
	}

	cur, err := c.customerCollection.Find(ctx, filter, findOptions)
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
