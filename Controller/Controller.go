package Controller

import (
	"Anton/Catalog"
	"Anton/View"
	"Anton/utils"
	"net/http"
	"strconv"
)

func AddFormController(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(View.AddPageView(View.CreateProductForm{}, "Добавьте новый продукт", "Добавить", "")))
}

func AddProductController(catalog Catalog.Catalog) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		name := r.FormValue("First")
		count, countErr := strconv.ParseInt(r.FormValue("Second"), 10, 64)
		price, priceErr := strconv.ParseFloat(r.FormValue("Third"), 64)

		hasError, createProductForm := utils.CheckError(*r, name, countErr, priceErr)

		if hasError {

			_, _ = w.Write([]byte(View.AddPageView(*createProductForm, "Добавьте новый продукт", "Попробовать снова", "")))

		} else {

			product, err := utils.CreateNewProduct(0, name, count, price)

			if err != nil {

			} else {
				_, _ = catalog.AddNewProduct(product)
				w.Header().Set("Location", "http://localhost:8080/list")
				w.WriteHeader(302)

			}
		}
	}
}

func EditProductController(catalog Catalog.Catalog) func(http.ResponseWriter, *http.Request) {

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
			_, _ = catalog.EditProduct(&Catalog.Product{id, name, count, price}) //!!!!!!!!

			http.Redirect(w, r, "http://localhost:8080/list", http.StatusFound)
		}
	}
}

func PrintListController(catalog Catalog.Catalog) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b := View.GeneratePageListHTMLController(catalog)
		_, _ = w.Write([]byte(View.ProductListView(b)))
	}
}

func ReturnToHomeController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://localhost:8080/add")
	w.WriteHeader(302)
}

//Эта функция выводит изменяемый объект
func FetchProductController(catalog Catalog.Catalog) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cameId, _ := strconv.Atoi(r.URL.Query().Get("product_id"))

		product, _ := catalog.GetProductByID(cameId)
		productForm := View.CreateProductForm{}
		productForm.name = product.name
		productForm.count = strconv.Itoa(int(product.count))
		productForm.price = strconv.Itoa(int(product.price))

		_, _ = w.Write([]byte(View.EditPageView(productForm, "Измените продукт", "Изменить", product.id)))
	}
}

func DeleteProductController(catalog Catalog.Catalog) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("product_id"))

		_ = catalog.DeleteProductById(id)

		http.Redirect(w, r, "http://localhost:8080/list", http.StatusFound)
	}
}
