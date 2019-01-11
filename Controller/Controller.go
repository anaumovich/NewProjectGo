package Controller

import (
	"Anton/CatalogModel"
	"Anton/View"
	"Anton/utils"
	"fmt"
	"net/http"
	"strconv"
)

func AddFormController(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(View.AddPageView(View.CreateProductForm{}, "Добавьте новый продукт", "Добавить", "")))
}

func AddProductController(catalog CatalogModel.Catalog) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		name := r.FormValue("First")
		count, countErr := strconv.ParseInt(r.FormValue("Second"), 10, 64)
		price, priceErr := strconv.ParseFloat(r.FormValue("Third"), 64)
		productType := r.FormValue("productType")

		hasError, createProductForm := utils.CheckError(*r, name, countErr, priceErr)

		if hasError {

			_, _ = w.Write([]byte(View.AddPageView(*createProductForm, "Добавьте новый продукт", "Попробовать снова", "")))

		} else {

			product, err := catalog.CreateNewProduct(0, name, count, price, productType)

			if err != nil {

			} else {
				_, _ = catalog.AddNewProduct(product)
				w.Header().Set("Location", "http://localhost:8080/list")
				w.WriteHeader(302)

			}
		}
	}
}

func EditProductController(catalog CatalogModel.Catalog) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		name := r.FormValue("First")
		count, countErr := strconv.ParseInt(r.FormValue("Second"), 10, 64)
		price, priceErr := strconv.ParseFloat(r.FormValue("Third"), 64)

		hasError, createProductForm := utils.CheckError(*r, name, countErr, priceErr)

		id, _ := strconv.Atoi(r.FormValue("product_id"))

		if hasError {
			if id != 0 {
				_, _ = w.Write([]byte(View.EditPageView(*createProductForm, "Измените продукт", "Изменить", id)))
			}
		} else {
			_, _ = catalog.EditProduct(id, name, count, price) //!!!!!!!!

			http.Redirect(w, r, "http://localhost:8080/list", http.StatusFound)
		}
	}
}

func PrintListController(catalog CatalogModel.Catalog) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(View.PrintProductList(catalog)))
	}
}

func ReturnToHomeController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://localhost:8080/add")
	w.WriteHeader(302)
}

func FetchProductController(catalog CatalogModel.Catalog) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cameId, _ := strconv.Atoi(r.URL.Query().Get("product_id"))

		product, _ := catalog.GetProductByID(cameId)

		productForm := View.CreateProductForm{}
		productForm.Name = product.GetName()
		productForm.Count = strconv.Itoa(int(product.GetCount()))
		productForm.Price = strconv.Itoa(int(product.GetPrice()))

		_, _ = w.Write([]byte(View.EditPageView(productForm, "Измените продукт", "Изменить", cameId)))
	}
}

func DeleteProductController(catalog CatalogModel.Catalog) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("product_id"))

		_ = catalog.DeleteProductById(id)

		http.Redirect(w, r, "http://localhost:8080/list", http.StatusFound)
	}
}

func SetDiscountController(catalog CatalogModel.Catalog) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		productType := r.FormValue("discountType")
		discount := r.FormValue("discount")
		fmt.Println(productType, discount)
		all := catalog.GetAll()
		for range all {
			fmt.Println()
		}

		http.Redirect(w, r, "http://localhost:8080/list", http.StatusFound)
	}
}
