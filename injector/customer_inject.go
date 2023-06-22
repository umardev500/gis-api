package injector

import (
	"gis/app/delivery"
	"gis/app/repository"
	"gis/app/service"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCustomerInjector(router fiber.Router, db *mongo.Database) {
	customerRepo := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepo)
	delivery.NewCustomerDelivery(router, customerService)
}
