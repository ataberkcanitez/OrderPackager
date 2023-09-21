package order

import (
	"github.com/ataberkcanitez/OrderPackager/internal/pack"
	"sort"
)

type OrderServiceImpl struct {
	PackService pack.PackService
}

func (os *OrderServiceImpl) CalculatePacksForOrder(itemsToShip int) ([]int, error) {
	packs, err := os.PackService.GetAllPacks()
	if err != nil {
		return nil, err
	}

	sortPacksBySizeDescending(packs)

	packsCounts := make([]int, len(packs))
	itemsRemaining := itemsToShip

	for i, pack := range packs {
		packsCounts[i] = itemsRemaining / pack.Size
		itemsRemaining %= pack.Size
	}

	return packsCounts, nil
}

func sortPacksBySizeDescending(packs []pack.Pack) {
	sortBySize := func(i, j int) bool {
		return packs[i].Size > packs[j].Size
	}
	sort.SliceStable(packs, sortBySize)
}
