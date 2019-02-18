package main

import (
	"NewProjectGo/CatalogModel"
	"NewProjectGo/Controller"
	"fmt"
	"net/http"
	"strings"
)

func controller(key string, catalog CatalogModel.Catalog) func(w http.ResponseWriter, r *http.Request) {
	switch key {
	case "GET/":
		return Controller.StartController()
	case "GET/login":
		return Controller.LoginController()
	case "GET/registration":
		return Controller.RegistrationController()
	case "GET/list":
		return Controller.PrintListController(catalog)
	case "GET/add":
		return Controller.AddFormController
	case "POST/add":
		return Controller.AddProductController(catalog)
	case "POST/list":
		return Controller.ReturnToHomeController
	case "GET/edit":
		return Controller.FetchProductController(catalog)
	case "POST/edit":
		return Controller.EditProductController(catalog)
	case "GET/delete":
		return Controller.DeleteProductController(catalog)

	default:
		return Controller.ErrorController()

	}
}

func handler(catalog CatalogModel.Catalog) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := strings.Join([]string{r.Method, r.URL.Path}, "")
		//fmt.Println(key)
		x := controller(key, catalog)
		x(w, r)
	}
}

func main() {

	Catalog := CatalogModel.CatalogConfigurator()

	http.HandleFunc("/", handler(Catalog))

	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err)
}
