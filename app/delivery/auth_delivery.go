package delivery

import "github.com/gofiber/fiber/v2"

type authDelivery struct{}

func NewAuthDelivery(router fiber.Router) {
	handler := &authDelivery{}
	auth := router.Group("/auth")

	auth.Post("/login", handler.Login)
}

func (a *authDelivery) Login(ctx *fiber.Ctx) error {
	return nil
}
