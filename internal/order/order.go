package order

type Order struct {
	ID          int
	ItemsToShip int
}

type OrderService interface {
	CalculatePacksForOrder(itemsToShip int) (map[int]int, error)
}
