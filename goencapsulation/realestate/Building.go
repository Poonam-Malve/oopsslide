package realestate

type Building struct {
	name    string
	address string
	size    string
}

func CreateBuilding(name, address, size string) *Building {
	return &Building{
		name:    name,
		address: address,
		size:    size,
	}
}

func (b *Building) GetName() string {
	return "Building name is: " + b.name
}
