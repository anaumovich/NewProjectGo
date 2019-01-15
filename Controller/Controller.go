package Controller

import (
	"NewProjectGo/CatalogModel"
	"NewProjectGo/utils"
	"NewProjectGo/view"
	"fmt"
	"net/http"
	"strconv"
)

func AddFormController(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte(view.AddPageView(view.CreateProductForm{}, "Добавьте новый продукт", "Добавить")))
}

func AddProductController(catalog CatalogModel.Catalog) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		name := r.FormValue("First")
		count, countErr := strconv.ParseInt(r.FormValue("Second"), 10, 64)
		price, priceErr := strconv.ParseFloat(r.FormValue("Third"), 64)
		productType := r.FormValue("productType")

		hasError, createProductForm := utils.CheckInputError(*r, name, countErr, priceErr)

		if hasError {

			_, _ = w.Write([]byte(view.AddPageView(*createProductForm, "Добавьте новый продукт", "Попробовать снова")))

		} else {
			product, err := CatalogModel.CreateNewProduct(name, count, price, productType)

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

		hasError, createProductForm := utils.CheckInputError(*r, name, countErr, priceErr)

		id, _ := strconv.Atoi(r.FormValue("product_id"))

		if hasError {
			if id != 0 {
				_, _ = w.Write([]byte(view.EditPageView(*createProductForm, "Изменить", id)))
			}
		} else {
			_, _ = catalog.EditProduct(id, name, count, price) //!!!!!!!!

			http.Redirect(w, r, "http://localhost:8080/list", http.StatusFound)
		}
	}
}

func PrintListController(catalog CatalogModel.Catalog) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(view.PrintProductList(catalog)))
		if err != nil {
			fmt.Println("cannot print product list")
		}
	}
}

func ReturnToHomeController(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Location", "http://localhost:8080/add")
	w.WriteHeader(302)
}

func FetchProductController(catalog CatalogModel.Catalog) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cameId, err := strconv.Atoi(r.URL.Query().Get("product_id"))
		if err != nil {
			fmt.Println("cannot get product id")
		}
		product, err := catalog.GetProductByID(cameId)
		if err != nil {
			fmt.Println("cannot get product bt id")
		}
		productForm := view.CreateProductForm{}
		productForm.Name = product.GetName()
		productForm.Count = strconv.Itoa(int(product.GetCount()))
		productForm.Price = strconv.Itoa(int(product.GetPrice()))

		_, err = w.Write([]byte(view.EditPageView(productForm, "Изменить", cameId)))
		if err != nil {
			fmt.Println("cannot edit product")
		}
	}
}

func DeleteProductController(catalog CatalogModel.Catalog) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("product_id"))
		if err != nil {
			fmt.Println("cannot get product id")
		}
		err = catalog.DeleteProductById(id)
		if err != nil {
			fmt.Println("cannot delete product by id")
		}
		http.Redirect(w, r, "http://localhost:8080/list", http.StatusFound)
	}
}

/*
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
*/
