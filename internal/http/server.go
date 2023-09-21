package http

import (
	"github.com/ataberkcanitez/OrderPackager/internal/order"
	"github.com/ataberkcanitez/OrderPackager/internal/pack"
	"github.com/gofiber/fiber/v2"
)

type HTTPServer struct {
	app          *fiber.App
	orderService order.OrderService
	packService  pack.PackService
}

func NewHTTPServer(app *fiber.App, orderService order.OrderService, packService pack.PackService) *HTTPServer {
	return &HTTPServer{
		app:          app,
		orderService: orderService,
		packService:  packService,
	}
}

func (s *HTTPServer) SetupRoutes() {
	s.app.Post("/calculate-packs", s.calculatePacksHandler)
}

func (s *HTTPServer) calculatePacksHandler(c *fiber.Ctx) error {
	var request struct {
		ItemsToShip int `json:"itemsToShip"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	packCounts, err := s.orderService.CalculatePacksForOrder(request.ItemsToShip)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{"packCounts": packCounts})
}
