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

type product struct {
	id                int
	name, productType string
	count             int64
	price             float64
}

func (p product) GetId() int {
	id := p.id
	return id
}
func (p product) GetName() string {
	name := p.name
	return name
}
func (p product) GetCount() int64 {
	count := p.count
	return count
}
func (p product) GetPrice() float64 {
	price := p.price
	return price
}

type Catalog interface {
	AddNewProduct(*product) (int, error)

	DeleteProductById(int) error

	GetAll() map[int]*product

	EditProduct(int, string, int64, float64) (int, error)

	GetProductByID(int) (*product, error)
}

func CreateNewProduct(name string, count int64, price float64, productType string) (*product, error) {
	if name == "" || count < 0 || price < 0 {
		return nil, errors.New("invalid product data")
	} else {
		Product := product{}

		Product.name = name
		Product.count = count
		Product.price = price
		Product.productType = productType

		return &Product, nil
	}
}

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
