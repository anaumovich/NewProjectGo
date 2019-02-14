package CatalogModel

import (
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

// todo make public Product
// todo use constructor for new inst creation move logic of validation to constructor
// todo Rule: new product don't has id. New it's - mean not saved.
// todo when user call catalog.AddNewProduct() catalog set id to product
// todo remove CreateNewProduct method from Catalog Interface and use Product construct.
// todo create fn NewProduct(name string, count int64, price float64, productType string)
// todo use pointer
// todo use AbstractFactory
// todo edit View
//

type Catalog interface {
	AddNewProduct(*Product) (int, error)

	DeleteProductById(int) error

	GetAll() (map[int]*Product, error)

	EditProduct(int, string, int64, float64) (int, error)

	GetProductByID(int) (*Product, error)
}

type CatalogFactory interface {
	CreateCatalog() Catalog
}

func CatalogConfigurator() Catalog {
	var catalog Catalog

	useFile := os.Args[1]

	if useFile == "f" {
		catalog = NewFileCatalogFactory().CreateCatalog()
		fmt.Println("localhost started with FileCatalog")
		return catalog
	}
	if useFile == "m" {
		catalog = NewInMemoryCatalogFactory().CreateCatalog()
		fmt.Println("localhost started with InMemoryCatalog")
		return catalog
	}
	if useFile == "db" {
		catalog = NewDBCatalogFactory().CreateCatalog()
		fmt.Println("localhost started with DbCatalog")
		return catalog
	}
	return catalog
}
