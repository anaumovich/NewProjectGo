package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

	catalog := SetCatalogType()

	handler := createRootHandler(getRoutes(&catalog))

	http.HandleFunc("/", http.HandlerFunc(handler))

	_ = http.ListenAndServe(":8080", nil)
}

func getRoutes(catalog *Catalog) (m map[string]func(w http.ResponseWriter, r *http.Request)) {
	m = make(map[string]func(w http.ResponseWriter, r *http.Request))
	m["GET/list"] = PrintListController(*catalog)
	m["GET/add"] = AddFormController
	m["POST/add"] = AddProductController(*catalog)
	m["POST/list"] = ReturnToHomeController
	m["GET/edit"] = FetchProductController(*catalog)
	m["POST/edit"] = EditProductController(*catalog)
	m["GET/delete"] = DeleteProductController(*catalog)
	return m
}

func createRootHandler(m map[string]func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := strings.Join([]string{r.Method, r.URL.Path}, "")
		fmt.Println(key)
		if result, ok := m[key]; ok == true {
			result(w, r)
		} else {
			w.Header().Set("Location", "http://localhost:8080/list")
			w.WriteHeader(302)
		}
	}
}
