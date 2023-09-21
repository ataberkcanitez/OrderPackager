package pack

type Pack struct {
	ID   int
	Size int
}

type PackService interface {
	GetAllPacks() ([]Pack, error)
	GetPackByID(id int) (*Pack, error)
}
