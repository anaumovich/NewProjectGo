package Anton

import (
	"Anton/Catalog"
	"Anton/Controller"
	"Anton/utils"
	"fmt"
	"net/http"
	"strings"
)

func main() {

	catalog := utils.SetCatalogType()

	handler := createRootHandler(getRoutes(&catalog))

	http.HandleFunc("/", http.HandlerFunc(handler))

	_ = http.ListenAndServe(":8080", nil)
}

func getRoutes(catalog *Catalog.Catalog) (m map[string]func(w http.ResponseWriter, r *http.Request)) {
	m = make(map[string]func(w http.ResponseWriter, r *http.Request))
	m["GET/list"] = Controller.PrintListController(*catalog)
	m["GET/add"] = Controller.AddFormController
	m["POST/add"] = Controller.AddProductController(*catalog)
	m["POST/list"] = Controller.ReturnToHomeController
	m["GET/edit"] = Controller.FetchProductController(*catalog)
	m["POST/edit"] = Controller.EditProductController(*catalog)
	m["GET/delete"] = Controller.DeleteProductController(*catalog)
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