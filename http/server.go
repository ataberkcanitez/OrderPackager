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
		Add(id string, pack *pack.Pack) error
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
	s.app.Post("/order/calculate-packs", s.calculatePacksHandler)
	s.app.Post("/packs", s.addPackHandler)
	s.app.Get("/packs", s.getAllPacksHandler)
	s.app.Get("/packs/:id", s.getPackByIDHandler)
}

func (s *HTTPServer) addPackHandler(c *fiber.Ctx) error {
	var request struct {
		ID   string `json:"id"`
		Size int    `json:"amount"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	pack := &pack.Pack{ID: request.ID, Size: request.Size}
	err := s.packService.Add(request.ID, pack)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error", "details": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true, "pack": pack})
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

	return c.JSON(fiber.Map{"packCounts": packCounts, "success": true})
}

func (s *HTTPServer) getAllPacksHandler(c *fiber.Ctx) error {
	allPacks, err := s.packService.GetAllPacks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.JSON(fiber.Map{"packs": allPacks, "success": true})
}

func (s *HTTPServer) getPackByIDHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	pack, err := s.packService.GetPackByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.JSON(fiber.Map{"pack": pack, "success": true})

}
