package CatalogModel

import (
	"github.com/pkg/errors"
)

type InMemoryCatalog struct {
	products map[int]*Product
}

type InMemoryCatalogFactory struct {
}

func NewInMemoryCatalogFactory() InMemoryCatalogFactory {
	InMemoryCatalogFactory := InMemoryCatalogFactory{}

	return InMemoryCatalogFactory
}

func (InMemoryCatalogFactory) CreateCatalog() Catalog {
	catalog := InMemoryCatalog{}
	catalog.products = make(map[int]*Product)
	return &catalog
}

//
func (catalog *InMemoryCatalog) AddNewProduct(product *Product) (int, error) {
	product.id = len(catalog.products) + 1
	catalog.products[product.id] = product

	return 0, errors.New("cannot add product")
}

func (catalog *InMemoryCatalog) DeleteProductById(cameId int) error {
	for key := cameId; key < len(catalog.products); key++ {
		catalog.products[key] = catalog.products[key+1]
		catalog.products[key].id--
	}
	delete(catalog.products, len(catalog.products))

	return errors.New("can't edit product")
}

func (catalog *InMemoryCatalog) GetAll() map[int]*Product {
	return catalog.products
}

func (catalog *InMemoryCatalog) EditProduct(cameId int, name string, count int64, price float64) (int, error) {
	var err error
	catalog.products[cameId].id = cameId
	catalog.products[cameId].name = name
	catalog.products[cameId].count = count
	catalog.products[cameId].price = price
    err = {}
	//здесь должна быть строка которая пересчитает размер скидки с учетом изменения стоимости
	return cameId, errors.Wrap(err,"can't edit product")
}

func (catalog *InMemoryCatalog) GetProductByID(cameId int) (*Product, error) {
	var err error
	if catalog.products[cameId] == nil {
		err = errors.New("no ID")
		}
	return catalog.products[cameId], err
}
