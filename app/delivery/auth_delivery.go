package delivery

import (
	"context"
	"gis/domain/model"
	"gis/domain/service"
	"gis/helper"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type authDelivery struct {
	service service.AuthServiceInterface
}

func NewAuthDelivery(router fiber.Router, service service.AuthServiceInterface) {
	handler := &authDelivery{
		service: service,
	}

	auth := router.Group("/auth")

	auth.Post("/login", handler.Login)
}

func (a *authDelivery) Login(ctx *fiber.Ctx) error {
	contx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var response model.Response
	var authCreds model.AuthCreds

	if err := ctx.BodyParser(&authCreds); err != nil {
		response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		return helper.APIResponse(ctx, fiber.StatusInternalServerError, response)
	}

	user, err := a.service.Login(contx, authCreds)
	if err != nil {
		// if error is not documents found
		if err == mongo.ErrNoDocuments {
			response = model.Response{
				Success: true,
				Status:  fiber.StatusNotFound,
				Message: "user not found",
			}
			return ctx.Status(fiber.StatusNotFound).JSON(response)
		}
		// response for exact error
		response = model.Response{
			Success: false,
			Status:  fiber.StatusInternalServerError,
			Error:   err.Error(),
		}
		return helper.APIResponse(ctx, fiber.StatusInternalServerError, response)
	}

	// create token
	t, err := helper.CreateToken(jwt.MapClaims{
		"exp": time.Now().Add(1 * time.Hour).Unix(),
	})

	if err != nil {

		response = model.Response{
			Success: false,
			Status:  fiber.StatusInternalServerError,
			Error:   err.Error(),
		}
		return helper.APIResponse(ctx, fiber.StatusInternalServerError, response)
	}

	cookie := fiber.Cookie{
		Name:     "token",
		Value:    *t,
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	data := map[string]interface{}{
		"user":  user,
		"token": t,
	}

	return ctx.JSON(model.Response{
		Success: true,
		Status:  fiber.StatusOK,
		Message: "login success",
		Data:    data,
	})
}
