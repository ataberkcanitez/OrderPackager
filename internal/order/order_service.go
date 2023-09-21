package order

import (
	"github.com/ataberkcanitez/OrderPackager/internal/pack"
	"sort"
)

type OrderServiceImpl struct {
	PackService pack.PackService
}

func (os *OrderServiceImpl) CalculatePacksForOrder(itemsToShip int) (map[int]int, error) {
	packs, err := os.PackService.GetAllPacks()
	if err != nil {
		return nil, err
	}

	sortPacksBySizeDescending(packs)

	packsCounts := make(map[int]int)
	itemsRemaining := itemsToShip

	for _, pack := range packs {
		packCount := itemsRemaining / pack.Size
		itemsRemaining %= pack.Size
		if packCount > 0 {
			packsCounts[pack.Size] = packCount
		}
	}

	return packsCounts, nil
}

func sortPacksBySizeDescending(packs []pack.Pack) {
	sortBySize := func(i, j int) bool {
		return packs[i].Size > packs[j].Size
	}
	sort.SliceStable(packs, sortBySize)
}
