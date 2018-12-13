package main

import (
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	name  string
	count int64
	price float64
}

type Catalog struct {
	products map[int]Product
	Add      func()
	Edit     func()
}

func main() {
	m := Catalog{map[int]Product{}, Add, Edit}
	m.Add()

}

func Add() {
	ww
}
func Edit() {
	fes
}
func my() {
	Catalog.Add()
}

//Эта функция выводит стартовую страницу
func startPage(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(StartPage()))
}

// Эта функция получает из POST запроса значения которые ввел пользователь на главной странице  записывает их в
// переменные и вызывает функцию Add которая добавляет новую запись в Store
// потом делает итерацию на следующий запись в Stor-e
// потом выводит новую страницу которая содержит все записи Stor-a
func GetList(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("First")
	count, _ := strconv.ParseInt(r.FormValue("Second"), 10, 64)
	price, _ := strconv.ParseFloat(r.FormValue("Third"), 64)
	Store[i] = Product{name, count, price}
	i++
	w.Header().Set("Location", "http://localhost:8080/add")
	_, _ = w.Write([]byte(NewPage(addString())))
}

//Эта функция возвращает на стартовую страницу при нажатии на кнопку назад кнопка назад прописана в
// form action и возвращает ссылку по которой обработчик запускает эту функцию
func ReturnToHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://localhost:8080")
	w.WriteHeader(302)
}

//Здесь из query string отправленной браузером получаем id
// редактируемого продукта и передаем этот id  странице редактирования продукта и ее выводим.
func EditData(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("product_id"))
	_, _ = w.Write([]byte(EditProduct(id)))
}

// Эта функция получает данные отправленные браузером методом GET а конкретно она получает значения отредактированного
// продукта и его ID, обновляет этот продукт в Store и выводит страницу с товарами включая изменения
func AddData(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("product_id"))
	editName := r.URL.Query().Get("One")
	editCount, _ := strconv.ParseInt(r.URL.Query().Get("Two"), 10, 64)
	editPrice, _ := strconv.ParseFloat(r.URL.Query().Get("Three"), 64)
	Store[id] = Product{editName, editCount, editPrice}
	_, _ = w.Write([]byte(NewPage(addString())))
}

//Эта функция добавляет новую запись на страницу с товарами
func addString() string {

	b := ""

	for a := 1; a < i; a++ {
		arr := make([]string, 8)
		arr[0] = `<tr><td>`
		arr[1] = string(Store[a].name)
		arr[2] = `</td><td>`
		arr[3] = strconv.FormatInt(Store[a].count, 10)
		arr[4] = `</td><td>`
		arr[5] = strconv.FormatFloat(Store[a].price, 'f', 0, 64)
		arr[6] = `</td><td><a href="http://localhost:8080/edit?product_id=` + strconv.Itoa(a) + `"><button>Изменить</button></a></td>`
		arr[7] = `</td></tr>`
		b += strings.Join(arr, "")
	}
	return b
}
