package main

import (
	"github.com/ataberkcanitez/order-packager/db"
	"github.com/ataberkcanitez/order-packager/http"
	"github.com/ataberkcanitez/order-packager/order"
	"github.com/ataberkcanitez/order-packager/pack"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	packMemDb := db.NewInMemDB[*pack.Pack]()
	initializePacks(packMemDb)

	packService := pack.NewPackService(packMemDb)
	orderService := order.NewOrderService(packService)

	httpServer := http.NewHTTPServer(app, orderService, packService)
	httpServer.SetupRoutes()

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}

func initializePacks(memDb *db.InMemDB[*pack.Pack]) {
	for _, p := range initialPacks {
		memDb.Save(p.ID, p)
	}
}

var initialPacks = []*pack.Pack{
	{ID: "1", Size: 250},
	{ID: "2", Size: 500},
	{ID: "3", Size: 1000},
	{ID: "4", Size: 2000},
	{ID: "5", Size: 5000},
}
