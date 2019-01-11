package utils

import (
	"Anton/CatalogModel"
	"Anton/View"
	"fmt"
	"net/http"
	"os"
)

// todo use AbstractFactory
func SetCatalogType() CatalogModel.Catalog {
	var catalog CatalogModel.Catalog

	useFile := os.Args[1]

	if useFile == "f" {
		catalog = CatalogModel.NewFileCatalog()
		fmt.Println("localhost started with FileCatalog")
		return catalog
	}
	if useFile == "m" {
		catalog = CatalogModel.NewInMemoryCatalog()
		fmt.Println("localhost started with InMemoryCatalog")
		return catalog
	}
	if useFile == "db" {
		catalog = CatalogModel.NewDbCatalog()
		fmt.Println("localhost started with DbCatalog")
		return catalog
	}
	return catalog
}

//

func CheckError(r http.Request, name string, countErr, priceErr error) (hasError bool, form *View.CreateProductForm) {

	hasError = false

	createProductForm := View.CreateProductForm{}

	createProductForm.Name = name
	createProductForm.Count = r.FormValue("Second")
	createProductForm.Price = r.FormValue("Third")

	if name == "" {
		createProductForm.NameError = "Ошибка имени"

		hasError = true
	}

	if countErr != nil {
		createProductForm.CountError = "Ошибка колличества"
		createProductForm.Count = r.FormValue("Second")

		hasError = true
	}

	if priceErr != nil {
		createProductForm.PriceError = "Ошибка стоимости"
		createProductForm.Price = r.FormValue("Third")

		hasError = true
	}

	return hasError, &createProductForm
}
