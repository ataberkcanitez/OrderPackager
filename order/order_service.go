package order

import (
	"github.com/ataberkcanitez/order-packager/pack"
	"sort"
)

type packService interface {
	GetAllPacks() ([]*pack.Pack, error)
}

type orderService struct {
	packSvc packService
}

func NewOrderService(packSvc packService) *orderService {
	return &orderService{
		packSvc: packSvc,
	}
}

func (os *orderService) CalculatePacksForOrder(itemsToShip int) ([]*OrderResponse, error) {
	packs, err := os.packSvc.GetAllPacks()
	if err != nil {
		return nil, err
	}

	sortPacksBySizeDescending(packs)

	orderResponse := []*OrderResponse{}
	itemsRemaining := itemsToShip

	for _, pack := range packs {
		packCount := itemsRemaining / pack.Size
		itemsRemaining %= pack.Size
		if packCount > 0 {
			orderResp := &OrderResponse{
				Pack:   pack,
				Amount: packCount,
			}
			orderResponse = append(orderResponse, orderResp)
		}
	}

	return orderResponse, nil
}

func sortPacksBySizeDescending(packs []*pack.Pack) {
	sortBySize := func(i, j int) bool {
		return packs[i].Size > packs[j].Size
	}
	sort.SliceStable(packs, sortBySize)
}
