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
	m["GET/add"] = AddFormController
	m["POST/add"] = AddProductController(*catalog)
	m["GET/list"] = PrintListController(*catalog)
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

/*
func ReadLine(catalog *FileCatalog, reader *bufio.Reader) {

	for {

		line, _, _ := reader.ReadLine()

		if len(line) == 0 {
			break
		}

		i := strings.IndexAny(string(line), "|")
		n := strings.IndexAny(string(line), "$")
		c := strings.IndexAny(string(line), "&")

		id, _ := strconv.Atoi(string(line[1:i]))
		name := string(line[i+1 : n])
		count, _ := strconv.ParseInt(string(line[n+1:c]), 10, 64)
		price, _ := strconv.ParseFloat(string(line[c+1:]), 64)

		//fmt.Println(id, name, count, price)
		/*Здесь можно будет добваить логику поиска ID

		_, _ = catalog.RestoreFromFile(&Product{name, count, price, id}, os.File{})
	}
}*/
/*
func FindPositionAndChangeLineInFile(checkId int, name string, count int64, price float64, catalog *FileCatalog, reader *bufio.Reader, file os.File) {

	_, _ = file.Seek(0, 0)
	a := name
	b := strconv.Itoa(int(count))
	c := strconv.Itoa(int(price))
	d := strconv.Itoa(checkId)

	currentDisplacement := 0

	for {

		line, _, _ := reader.ReadLine()
		lineX := string(line) + "\n"

		StartCurrentString := len(lineX)

		currentDisplacement += len(lineX)

		thisPosition := currentDisplacement - StartCurrentString

		if len(line) == 0 {
			break
		}

		id, _ := strconv.Atoi(string(line[1:strings.IndexAny(string(line), "|")]))

		if checkId == id {

			textImpression := "#" + d + "|" + a + "$" + b + "&" + c + "\n"

			_, _ = file.Seek(int64(thisPosition), 0) // устанавливаем курсор в позицию записи

			_, _ = file.WriteString(textImpression) //записываем введенную строку

			_ = file.Truncate(int64(thisPosition + len(textImpression))) // обрезаем файл по последнему байту этой строки

			for i := checkId + 1; i <= catalog.lastId; i++ { // в конец строки пушим из памяти все данные из виртуальной бд
				x := catalog.products[i].name
				y := strconv.Itoa(int(catalog.products[i].count))
				z := strconv.Itoa(int(catalog.products[i].price))
				w := strconv.Itoa(catalog.products[i].id)

				str := "#" + w + "|" + x + "$" + y + "&" + z + "\n"

				_, _ = file.WriteString(str)

			}
		}
	}
}*/
