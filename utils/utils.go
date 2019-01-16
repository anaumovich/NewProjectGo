package utils

import (
	"NewProjectGo/view"
	"net/http"
)

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
