package pack

import (
	"fmt"
	"github.com/ataberkcanitez/order-packager/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPackServiceImpl_GetAllPacks(t *testing.T) {
	packDB := db.NewInMemDB[*Pack]()
	prepareTestData(packDB)

	packService := NewPackService(packDB)
	packs, err := packService.GetAllPacks()
	assert.Nil(t, err)
	assert.Len(t, packs, 5)

	expectedSizes := []int{250, 500, 1000, 2000, 5000}
	for _, expectedSize := range expectedSizes {
		found := false
		for _, pack := range packs {
			if pack.Size == expectedSize {
				found = true
				break
			}
		}
		assert.True(t, found)
	}
}

func TestPackServiceImpl_GetPackByID(t *testing.T) {
	packDB := db.NewInMemDB[*Pack]()
	prepareTestData(packDB)
	packService := NewPackService(packDB)

	pack, err := packService.GetPackByID("3")
	assert.Nil(t, err)
	assert.Equal(t, 1000, pack.Size)

	pack, err = packService.GetPackByID("99")
	assert.NotNil(t, err)
	assert.Nil(t, pack)
	assert.Equal(t, "Pack not found", err.Error())
	fmt.Println(err.Error())
}

func prepareTestData(d *db.InMemDB[*Pack]) {
	for _, testPack := range testPacks {
		d.Save(testPack.ID, testPack)
	}
}

var testPacks = []*Pack{
	{ID: "1", Size: 250},
	{ID: "2", Size: 500},
	{ID: "3", Size: 1000},
	{ID: "4", Size: 2000},
	{ID: "5", Size: 5000},
}
