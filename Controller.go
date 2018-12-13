package main

import (
	"net/http"
	"strings"
)

var i = 1

var Store = make(map[int]Product)

func main() {
	handler := createHandler(getRoutes())

	http.HandleFunc("/", http.HandlerFunc(handler))
	_ = http.ListenAndServe(":8080", nil)

}

// Эта функция описывает вызовы функций в зависимочти от  типа запроса пришедшего от браузера
func getRoutes() (m map[string]func(w http.ResponseWriter, r *http.Request)) {
	m = make(map[string]func(w http.ResponseWriter, r *http.Request))
	m["GET/"] = startPage
	m["POST/list"] = GetList
	m["POST/redirect"] = ReturnToHome
	m["GET/edit"] = EditData // regexp
	m["GET/addProduct"] = AddData
	return m
}


func createHandler(m map[string]func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := strings.Join([]string{r.Method, r.URL.Path}, "")
		//fmt.Println(key)
		if result, ok := m[key]; ok == true {
			result(w, r)
		} else {
			w.WriteHeader(404)
		}
	}
}



