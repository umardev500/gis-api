package injector

import (
	"gis/app/delivery"
	"gis/app/repository"
	"gis/app/service"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAuthInjector(router fiber.Router, db *mongo.Database) {
	userRepo := repository.NewUserRepository(db)
	service := service.NewAuthService(userRepo)
	delivery.NewAuthDelivery(router, service)
}
