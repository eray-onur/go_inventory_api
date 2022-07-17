package entities

type Product struct {
	Base         BaseEntity
	Title        string
	Description  string
	UnitsInStock float64
	Price        float64
	Cost         float64
}

func (p *Product) CalculatePrice() float64 {
	return p.Price
}
