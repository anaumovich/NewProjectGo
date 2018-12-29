package utils

import (
	"Anton/Catalog"
	"Anton/View"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func SetCatalogType() Catalog.Catalog {

	var catalog Catalog.Catalog

	useFile := os.Args[1]

	if useFile == "f" {
		catalog = Catalog.NewFileCatalog()
		fmt.Println("localhost started with FileCatalog")
		return catalog
	}

	if useFile == "m" {
		catalog = Catalog.NewInMemoryCatalog()
		fmt.Println("localhost started with InMemoryCatalog")
		return catalog
	} else {
		SetCatalogType()
	}
	return catalog
}

//It's OK
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

func CreateNewProduct(id int, name string, count int64, price float64) (*Catalog.Product, error) {
	if name == "" || count < 0 || price < 0 {
		return nil, errors.New("invalid product data")
	} else {
		product := Catalog.Product{id, name, count, price}

		return &product, nil
	}
}

func CheckError(r http.Request, name string, countErr, priceErr error) (hasError bool, form *View.CreateProductForm) {

	hasError = false

	createProductForm := View.CreateProductForm{}

	createProductForm.name = name
	createProductForm.count = r.FormValue("Second")
	createProductForm.price = r.FormValue("Third")

	if name == "" {
		createProductForm.nameError = "Ошибка имени"

		hasError = true
	}

	if countErr != nil {
		createProductForm.countError = "Ошибка колличества"
		createProductForm.count = r.FormValue("Second")

		hasError = true
	}

	if priceErr != nil {
		createProductForm.priceError = "Ошибка стоимости"
		createProductForm.price = r.FormValue("Third")

		hasError = true
	}

	return hasError, &createProductForm
}
