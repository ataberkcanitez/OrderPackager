package order

import "github.com/ataberkcanitez/OrderPackager/internal/pack"

type defaultOrderService struct {
	packService pack.PackService
}

func NewOrderService(packService pack.PackService) OrderService {
	return &defaultOrderService{packService: packService}
}

func (o *defaultOrderService) CalculatePacksForOrder(order Order) []int {
	return nil
}
