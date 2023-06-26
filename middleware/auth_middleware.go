package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Authentication(ctx *fiber.Ctx) error {
	fmt.Println("middleware running")
	return ctx.Next()
}
