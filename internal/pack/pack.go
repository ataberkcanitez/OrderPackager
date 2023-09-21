package pack

type Pack struct {
	Size   int
	Amount int
}

type PackService interface {
	CalculatePacks(items int) []Pack
}
