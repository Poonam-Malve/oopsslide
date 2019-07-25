package realstate

type House struct {
	name    string
	address string
	rooms   int
}

func CreateHouse(name, address string, rooms int) *House {
	return &House{
		name:    name,
		address: address,
		rooms:   rooms,
	}
}

func (h *House) GetName() string {
	return "House name is: " + h.name
}
