package CatalogModel

import "errors"

type InMemoryCatalog struct {
	products map[int]*product
}

func NewInMemoryCatalog() *InMemoryCatalog {

	catalog := InMemoryCatalog{}
	catalog.products = make(map[int]*product)

	return &catalog
}

func (catalog InMemoryCatalog) AddNewProduct(product *product) (int, error) {

	a := len(catalog.products)
	product.id = a + 1
	catalog.products[a+1] = product

	return 0, errors.New("cannot add product")
}

func (catalog InMemoryCatalog) DeleteProductById(cameId int) error {
	for key := cameId; key < len(catalog.products); key++ {
		catalog.products[key] = catalog.products[key+1]
		catalog.products[key].id--
	}
	delete(catalog.products, len(catalog.products))

	return errors.New("can't edit product")
}

func (catalog InMemoryCatalog) GetAll() map[int]*product {
	return catalog.products
}

func (catalog InMemoryCatalog) EditProduct(cameId int, name string, count int64, price float64) (int, error) {

	catalog.products[cameId].id = cameId
	catalog.products[cameId].name = name
	catalog.products[cameId].count = count
	catalog.products[cameId].price = price

	//здесь должна быть строка которая пересчитает размер скидки с учетом изменения стоимости

	return cameId, errors.New("can't edit product")
}

func (catalog InMemoryCatalog) GetProductByID(cameId int) (*product, error) {
	return catalog.products[cameId], errors.New("product not found")
}
