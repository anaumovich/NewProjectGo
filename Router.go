package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {

	file := OpenOrCreateFile()
	reader := bufio.NewReader(file)

	Catalog := newCatalog()
	ReadLine(Catalog, reader)

	handler := createRootHandler(getRoutes(Catalog, *file, *reader))
	http.HandleFunc("/", http.HandlerFunc(handler))

	_ = http.ListenAndServe(":8080", nil)
}

func getRoutes(Catalog *Catalog, file os.File, reader bufio.Reader) (m map[string]func(w http.ResponseWriter, r *http.Request)) {
	m = make(map[string]func(w http.ResponseWriter, r *http.Request))
	m["GET/add"] = AddFormController
	m["POST/add"] = AddProductController(Catalog, file)
	m["GET/list"] = PrintListController(Catalog)
	m["POST/list"] = ReturnToHomeController
	m["GET/edit"] = FetchProductController(Catalog)
	m["POST/edit"] = EditProductController(Catalog, file, reader)
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

func OpenOrCreateFile() *os.File {

	_, err := os.Stat("MyFile.txt")
	if err != nil {
		file, _ := os.Create("MyFile.txt")
		fmt.Println("I create File")
		return file
	} else {
		file, _ := os.OpenFile("MyFile.txt", os.O_RDWR, 111)
		fmt.Println("I open File")
		return file
	}

}

func ReadLine(catalog *Catalog, reader *bufio.Reader) {
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

		/*Здесь можно будет добваить логику поиска ID*/

		_, _ = catalog.ReadProductsFromFileAndWriteThemInCatalog(&Product{name, count, price, id}, os.File{})
	}
}

func FindPositionAndChangeLineInFile(checkId int, name string, count int64, price float64, catalog *Catalog, reader *bufio.Reader, file os.File) {

	_, _ = file.Seek(0, 0)
	a := name
	b := strconv.Itoa(int(count))
	c := strconv.Itoa(int(price))
	d := strconv.Itoa(checkId)

	lenght := 0

	for {

		line, _, _ := reader.ReadLine()

		if len(line) == 0 {
			break
		}

		id, _ := strconv.Atoi(string(line[1:strings.IndexAny(string(line), "|")]))

		this := len(line)

		lenght += len(line)

		if checkId == id {

			textImpression := "\n" + "#" + d + "|" + a + "$" + b + "&" + c + "\n"

			if id == 1 {
				textImpression = "#" + d + "|" + a + "$" + b + "&" + c + "\n"
			}
			_, _ = file.Seek(0, 0)

			_, _ = file.Seek(int64(this), 1)

			_, _ = file.WriteString(textImpression)

			for i := checkId + 1; i <= catalog.lastId; i++ {
				a := catalog.products[i].name
				b := strconv.Itoa(int(catalog.products[i].count))
				c := strconv.Itoa(int(catalog.products[i].price))
				d := strconv.Itoa(catalog.products[i].id)

				str := "#" + d + "|" + a + "$" + b + "&" + c + "\n"

				_, _ = file.WriteString(str)

			}

		}

	}
}
