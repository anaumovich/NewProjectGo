package Controller

import (
	"NewProjectGo/CatalogModel"
	"NewProjectGo/Utils"
	"NewProjectGo/View"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func StartController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(View.PrintStartPage("Зарегистрируйтесь или авторизуйтесь в системе")))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ErrorController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(View.PrintDefaultPage("fatal error")))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func LoginController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(View.PrintLoginPage("Вход")))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func RegistrationController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(View.PrintDefaultPage("Регистрация в системе")))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func AddFormController(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte(View.AddPageView(View.CreateProductForm{}, "Добавьте новый продукт", "Добавить")))
	fmt.Println(errors.Wrap(err, "AddFormController:"))
}

func AddProductController(catalog CatalogModel.Catalog) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		name := r.FormValue("First")
		count, countErr := strconv.ParseInt(r.FormValue("Second"), 10, 64)
		price, priceErr := strconv.ParseFloat(r.FormValue("Third"), 64)
		productType := r.FormValue("productType")

		hasError, createProductForm := Utils.CheckInputError(*r, name, countErr, priceErr)

		if hasError {

			_, err := w.Write([]byte(View.AddPageView(*createProductForm, "Добавьте новый продукт", "Попробовать снова")))
			if err != nil {
				fmt.Println(errors.Wrap(err, "AddProductController"))
			}

		} else {
			product := CatalogModel.CreateNewProduct(name, productType, count, price)
			_, err := catalog.AddNewProduct(product)
			if err != nil {
				fmt.Println(errors.Wrap(err, "AddProductController"))

			}
			w.Header().Set("Location", "http://localhost:8080/list")
			w.WriteHeader(302)
		}
	}
}

func EditProductController(catalog CatalogModel.Catalog) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		name := r.FormValue("First")
		count, countErr := strconv.ParseInt(r.FormValue("Second"), 10, 64)
		price, priceErr := strconv.ParseFloat(r.FormValue("Third"), 64)

		hasError, createProductForm := Utils.CheckInputError(*r, name, countErr, priceErr)

		id, err := strconv.Atoi(r.FormValue("product_id"))
		if err != nil {
			fmt.Println(errors.Wrap(err, ""))
		}
		if hasError {
			if id != 0 {
				_, err = w.Write([]byte(View.EditPageView(*createProductForm, "Изменить", id)))
			}
		} else {
			_, err = catalog.EditProduct(id, name, count, price) //!!!!!!!!

			http.Redirect(w, r, "http://localhost:8080/list", http.StatusFound)
		}
	}
}

func PrintListController(catalog CatalogModel.Catalog) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(View.PrintProductList(catalog)))
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
		productForm := View.CreateProductForm{}
		productForm.Name = product.GetName()
		productForm.Count = strconv.Itoa(int(product.GetCount()))
		productForm.Price = strconv.Itoa(int(product.GetPrice()))

		_, err = w.Write([]byte(View.EditPageView(productForm, "Изменить", cameId)))
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
