package pack

import "errors"

type packDB interface {
	Get(key string) (*Pack, bool)
	GetAll() ([]*Pack, error)
	Save(key string, val *Pack)
}

type packService struct {
	db packDB
}

func NewPackService(db packDB) *packService {
	return &packService{db}
}

func (ps *packService) GetAllPacks() ([]*Pack, error) {
	return ps.db.GetAll()
}

var ErrPackNotFound = errors.New("Pack not found")

func (ps *packService) GetPackByID(id string) (*Pack, error) {
	pack, ok := ps.db.Get(id)
	if !ok {
		return nil, ErrPackNotFound
	}
	return pack, nil
}
