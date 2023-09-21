package http

import (
	"github.com/ataberkcanitez/order-packager/order"
	"github.com/ataberkcanitez/order-packager/pack"
	"github.com/gofiber/fiber/v2"
)

type (
	orderService interface {
		CalculatePacksForOrder(itemsTShip int) ([]*order.OrderResponse, error)
	}

	packService interface {
		GetAllPacks() ([]*pack.Pack, error)
		GetPackByID(id string) (*pack.Pack, error)
	}
)

type HTTPServer struct {
	app          *fiber.App
	orderService orderService
	packService  packService
}

func NewHTTPServer(app *fiber.App, orderService orderService, packService packService) *HTTPServer {
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
