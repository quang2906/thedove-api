package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"ocg.com/database"
	"ocg.com/routes"
)

func main() {
	database.Connect()

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Use(cors.New())

	productRouter := app.Group("/api/v1/products")
	routes.ConfigProductRouter(&productRouter)

	categoryRouter := app.Group("/api/v1/categories")
	routes.ConfigCategoryRouter(&categoryRouter)

	userRouter := app.Group("/api/v1/users")
	routes.ConfigUserRouter(&userRouter)

	// doan vue dau?

	reviewRouter := app.Group("/api/v1/reviews")
	routes.ConfigReviewRouter(&reviewRouter)

	app.Listen(":3000")

}
