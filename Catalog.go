package main

import (
	"errors"
	"os"
	"strconv"
)

type Storage struct {
	// private
	products map[int]Product
}

func NewStorage() *Storage {
	return &Storage{products: make(map[int]Product)}
}

type Catalog struct {
	storage *Storage

	products map[int]*Product

	lastId int
}

// todo getProductWithId(id int) (*Product, error)

// todo Naming for constructors
//конструктор
func newCatalog() *Catalog {
	catalog := &Catalog{}

	catalog.storage = NewStorage()

	catalog.products = make(map[int]*Product)

	return catalog
}

type Product struct {
	name  string
	count int64

	price float64
	id    int
}

// todo add setters for

func (catalog *Catalog) AddNewProduct(product *Product, file os.File) (int, error) {
	if product.id != 0 {
		return 0, errors.New("product already have id")
	} else {
		catalog.lastId++
		product.id = catalog.lastId
		catalog.products[catalog.lastId] = product

		a := catalog.products[catalog.lastId].name
		b := strconv.Itoa(int(catalog.products[catalog.lastId].count))
		c := strconv.Itoa(int(catalog.products[catalog.lastId].price))
		d := strconv.Itoa(int(catalog.products[catalog.lastId].id))

		textImpression := "#" + d + "|" + a + "$" + b + "&" + c + "\n"
		_, _ = file.Write([]byte(textImpression))

	}

	return catalog.lastId, nil

}

func (catalog *Catalog) ReadProductsFromFileAndWriteThemInCatalog(product *Product, file os.File) (int, error) {
	catalog.lastId++
	product.id = catalog.lastId
	catalog.products[catalog.lastId] = product

	return catalog.lastId, nil
}

func createNewProduct(name string, count int64, price float64, id int) (*Product, error) {
	if name == "" || count < 0 || price < 0 {
		return nil, errors.New("invalid product data")
	} else {
		MyProduct := Product{name, count, price, id}

		return &MyProduct, nil
	}
}

//бизнес логика

func (*Product) CheckNameSetter(name string, id int) (bool, error) {
	if name == "" {
		return true, errors.New("invalid product name")
	}
	return false, nil
}

func (*Product) CheckCountSetter(count int64, id int) (bool, error) {
	if count < 0 {
		return true, errors.New("invalid product count")
	}
	return false, nil
}

func (*Product) CheckPriceSetter(price float64, id int) (bool, error) {
	if price <= 0 {
		return true, errors.New("invalid product price")
	}
	return false, nil
}
