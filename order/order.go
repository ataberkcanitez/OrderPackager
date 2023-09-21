package order

import (
	"github.com/ataberkcanitez/order-packager/pack"
)

type Order struct {
	ID          int
	ItemsToShip int
}

type OrderResponse struct {
	Pack   *pack.Pack
	Amount int
}
