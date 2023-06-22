package helper

import "github.com/gofiber/fiber/v2"

func APIResponse(ctx *fiber.Ctx, status int, data interface{}) error {
	return ctx.Status(status).JSON(data)
}
