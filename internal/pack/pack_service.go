package pack

import "sort"

type defaultPackService struct {
	packSizes []int
}

func NewPackService(packSizes []int) PackService {
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))
	return &defaultPackService{packSizes: packSizes}
}

func (p *defaultPackService) CalculatePacks(items int) []Pack {
	return nil
}
