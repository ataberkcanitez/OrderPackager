package order

import "github.com/ataberkcanitez/OrderPackager/internal/pack"

type Order struct {
	ID          int
	ItemsToShip int
}

type OrderResponse struct {
	Pack   pack.Pack
	Amount int
}

type OrderService interface {
	CalculatePacksForOrder(itemsToShip int) ([]*OrderResponse, error)
}
