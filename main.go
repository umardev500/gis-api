package main

import (
	"fmt"
	"gis/config"
	"gis/injector"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	db := config.NewConn().Database("db_gis")

	port := os.Getenv("PORT")
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("api")

	injector.NewCustomerInjector(api, db)
	injector.NewAuthInjector(api, db)

	log.Fatal(app.Listen(port))
}
