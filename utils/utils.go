package utils

import (
	"AmazingCatalog/CatalogModel"
	"AmazingCatalog/view"
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

func CheckInputError(r http.Request, name string, countErr, priceErr error) (hasError bool, form *view.CreateProductForm) {

	hasError = false

	createProductForm := view.CreateProductForm{}

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
