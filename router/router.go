package router

import (
	"app/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// version
	v1 := app.Group("/v1", logger.New())

	// Products
	product := v1.Group("/product")
	product.Get("/", controllers.GetAllProducts)
	product.Get("/:id", controllers.GetProduct)
	product.Post("/", controllers.CreateProduct)
	product.Put("/", controllers.UpdateProduct)
	product.Delete("/:id", controllers.DeleteProduct)
	product.Delete("/", controllers.DeleteProducts)
	// product.Delete("/", middlewares.Protected(), handlers.DeleteProduct) // jwt
}
