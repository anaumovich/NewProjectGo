package CatalogModel

import (
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

// todo make public
// todo use constructor for new inst creation move logic of validation to constructor

// todo Rule: new product don't has id. New it's - mean not saved.
// todo when user call catalog.AddNewProduct() catalog set id to product

type Product struct {
	id                int
	name, productType string
	count             int64
	price             float64
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

func CreateNewProduct(name string, count int64, price float64, productType string) (*Product, error) {
	if name == "" || count < 0 || price < 0 {
		return nil, errors.New("invalid product data")
	} else {
		product := Product{}

		product.name = name
		product.count = count
		product.price = price
		product.productType = productType

		return &product, nil
	}
}

//
//
type CatalogFactory interface {
	CreateCatalog() Catalog
}

type Catalog interface {
	AddNewProduct(*Product) (int, error)

	DeleteProductById(int) error

	GetAll() map[int]*Product

	EditProduct(int, string, int64, float64) (int, error)

	GetProductByID(int) (*Product, error)
}

//
func OpenOrCreateFile() *os.File {

	_, err := os.Stat("MyFile.txt")
	if err != nil {
		file, _ := os.Create("MyFile.txt")
		fmt.Println("I create File")
		return file
	} else {
		file, _ := os.OpenFile("MyFile.txt", os.O_RDWR, 111)
		fmt.Println("I open File")
		return file
	}
}
