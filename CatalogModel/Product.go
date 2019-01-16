package CatalogModel

type Product struct {
	id                int
	name, productType string
	count             int64
	price             float64
}

func CreateNewProduct(name string, productType string, count int64, price float64) *Product {
	return &Product{name: name, productType: productType, count: count, price: price}
}

func (p Product) GetId() int {
	id := p.id
	return id
}
func (p Product) GetName() string {
	name := p.name
	return name
}
func (p Product) GetCount() int64 {
	count := p.count
	return count
}
func (p Product) GetPrice() float64 {
	price := p.price
	return price
}
