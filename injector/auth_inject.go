package injector

import (
	"gis/app/delivery"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAuthInjector(router fiber.Router, db *mongo.Database) {
	delivery.NewAuthDelivery(router)
}
