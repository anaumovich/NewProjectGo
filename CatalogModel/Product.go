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

func (product Product) GetId() int {
	id := product.id
	return id
}
func (product Product) GetName() string {
	name := product.name
	return name
}
func (product Product) GetCount() int64 {
	count := product.count
	return count
}
func (product Product) GetPrice() float64 {
	price := product.price
	return price
}
