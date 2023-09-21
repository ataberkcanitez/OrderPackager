package pack

import "errors"

var packs = []Pack{
	{ID: 1, Size: 250},
	{ID: 2, Size: 500},
	{ID: 3, Size: 1000},
	{ID: 4, Size: 2000},
	{ID: 5, Size: 5000},
}

type PackServiceImpl struct{}

func (ps *PackServiceImpl) GetAllPacks() ([]Pack, error) {
	return packs, nil
}

func (ps *PackServiceImpl) GetPackByID(id int) (*Pack, error) {
	for _, pack := range packs {
		if pack.ID == id {
			return &pack, nil
		}
	}

	return nil, errors.New("Pack not found")
}
