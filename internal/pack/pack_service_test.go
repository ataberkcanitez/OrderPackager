package pack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPackServiceImpl_GetAllPacks(t *testing.T) {
	packService := &PackServiceImpl{}
	packs, err := packService.GetAllPacks()
	assert.Nil(t, err)
	assert.Len(t, packs, 5)

	expectedSizes := []int{250, 500, 1000, 2000, 5000}
	for i, pack := range packs {
		assert.Equal(t, expectedSizes[i], pack.Size)
	}
}

func TestPackServiceImpl_GetPackByID(t *testing.T) {
	packService := &PackServiceImpl{}

	pack, err := packService.GetPackByID(3)
	assert.Nil(t, err)
	assert.Equal(t, 1000, pack.Size)

	pack, err = packService.GetPackByID(99)
	assert.NotNil(t, err)
	assert.Nil(t, pack)
	assert.Equal(t, "Pack not found", err.Error())
	fmt.Println(err.Error())
}
