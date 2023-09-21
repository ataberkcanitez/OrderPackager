package order

type Order struct {
	Items int
}

type OrderService interface {
	CalculatePacksForOrder(order Order) []int
}
