package order

import (
	"github.com/ataberkcanitez/order-packager/pack"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderServiceImplementation_CalculatePacksForOrder(t *testing.T) {

	packService := &mockPackService{}
	orderService := NewOrderService(packService)

	// Test case: 12001 items to ship
	packCounts, err := orderService.CalculatePacksForOrder(12001)
	assert.Nil(t, err)

	expectedPackCounts := map[int]int{
		5000: 2,
		2000: 1,
	}

	for _, packCount := range packCounts {
		expected := expectedPackCounts[packCount.Pack.Size]
		assert.Equal(t, expected, packCount.Amount)

	}
}

type mockPackService struct{}

func (m *mockPackService) GetAllPacks() ([]*pack.Pack, error) {
	return []*pack.Pack{
		{ID: "1", Size: 250},
		{ID: "2", Size: 500},
		{ID: "3", Size: 1000},
		{ID: "4", Size: 2000},
		{ID: "5", Size: 5000},
	}, nil
}

func (m *mockPackService) GetPackByID(id string) (*pack.Pack, error) {
	for _, pack := range []pack.Pack{
		{ID: "1", Size: 250},
		{ID: "2", Size: 500},
		{ID: "3", Size: 1000},
		{ID: "4", Size: 2000},
		{ID: "5", Size: 5000},
	} {
		if pack.ID == id {
			return &pack, nil
		}
	}
	return nil, nil
}
