package main

import (
	"fmt"
	"github.com/ataberkcanitez/OrderPackager/internal/http"
	"github.com/ataberkcanitez/OrderPackager/internal/order"
	"github.com/ataberkcanitez/OrderPackager/internal/pack"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	packService := &pack.PackServiceImpl{}
	orderService := &order.OrderServiceImpl{PackService: packService}

	httpServer := http.NewHTTPServer(app, orderService, packService)
	httpServer.SetupRoutes()

	err := app.Listen(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
