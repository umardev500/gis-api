package delivery

import (
	"context"
	"fmt"
	"gis/domain/model"
	"gis/domain/service"
	"gis/helper"
	"gis/middleware"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type customerDelivery struct {
	service service.CustomerServiceInterface
}

func NewCustomerDelivery(router fiber.Router, service service.CustomerServiceInterface) {
	handler := &customerDelivery{
		service: service,
	}
	secret := os.Getenv("SECRET")

	router = router.Group("customer")
	router.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.JSON(model.Response{
				Success: false,
				Status:  fiber.StatusBadRequest,
				Error:   err.Error(),
			})
		},
	}))

	router.Post("/", handler.Create)
	router.Get("/", middleware.Authentication, handler.FindAll)
	// get nearest
	router.Get("/near", middleware.Authentication, handler.FindAllNearest)
	router.Get("/:id", handler.FindOne)
}

func (c *customerDelivery) Create(ctx *fiber.Ctx) error {
	contx, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var customer model.CustomerRequestPayload
	if err := ctx.BodyParser(&customer); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.Response{
			Success: false,
			Error:   "Failed to decode customer",
		})
	}

	fmt.Println("customer", customer)

	// call service
	err := c.service.Create(contx, customer)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Success: false,
			Error:   "Internal server error",
		})
	}

	return ctx.JSON(model.Response{
		Success: true,
		Message: "Customer successfuly added",
	})
}

func (c *customerDelivery) FindAll(ctx *fiber.Ctx) error {
	contx, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var response model.Response

	perPage, _ := strconv.Atoi(ctx.Query("per_page", "0"))
	order := ctx.Query("order")
	findMeta := model.FindMetaRequest{
		PerPage: int64(perPage),
		Order:   order,
	}
	customers, meta, err := c.service.FindAll(contx, findMeta)
	if err != nil {
		response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		return helper.APIResponse(ctx, fiber.StatusInternalServerError, response)
	}

	response = model.Response{
		Success: true,
		Message: "get all customers",
		Data:    customers,
		Meta:    &meta,
	}

	return ctx.JSON(response)
}

func (c *customerDelivery) FindAllNearest(ctx *fiber.Ctx) error {
	contx, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var response model.Response

	perPage, _ := strconv.Atoi(ctx.Query("per_page", "0"))
	order := ctx.Query("order")
	findMeta := model.FindMetaRequest{
		PerPage: int64(perPage),
		Order:   order,
	}
	customers, meta, err := c.service.FindAllNearest(contx, findMeta)
	if err != nil {
		response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		return helper.APIResponse(ctx, fiber.StatusInternalServerError, response)
	}

	response = model.Response{
		Success: true,
		Message: "get all customers",
		Data:    customers,
		Meta:    &meta,
	}

	return ctx.JSON(response)
}

func (c *customerDelivery) FindOne(ctx *fiber.Ctx) error {
	contx, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")

	var response model.Response

	customer, err := c.service.FindOne(contx, id)
	if err != nil {
		// if error is not documents found
		if err == mongo.ErrNoDocuments {
			response = model.Response{
				Success: true,
				Message: "get customer by id",
			}
			return ctx.JSON(response)
		}
		// response for exact error
		response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		return helper.APIResponse(ctx, fiber.StatusInternalServerError, response)
	}

	response = model.Response{
		Success: true,
		Message: "get customer by id",
		Data:    customer,
	}

	return ctx.JSON(response)
}
