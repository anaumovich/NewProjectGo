package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

/*import (
	"errors"
)

func newCatalog() *FileCatalog {

	catalog := FileCatalog{}

	catalog.products = make(map[int]*Product)

	//catalog.storage = NewStorage()

	return &catalog
}*/

/*
func AddNewProduct(product *Product, file os.File) (int, error) {
	if product.id != 0 {
		return 0, errors.New("product already have id")
	} else {
		catalog.lastId++
		product.id = catalog.lastId
		catalog.products[catalog.lastId] = product

		a := catalog.products[catalog.lastId].name
		b := strconv.Itoa(int(catalog.products[catalog.lastId].count))
		c := strconv.Itoa(int(catalog.products[catalog.lastId].price))
		d := strconv.Itoa(int(catalog.products[catalog.lastId].id))

		textImpression := "#" + d + "|" + a + "$" + b + "&" + c + "\n"
		_, _ = file.Write([]byte(textImpression))

	}

	return catalog.lastId, nil

}*/

// ReadProductsFromFileAndWriteThemInCatalog
/*func (catalog *FileCatalog) RestoreFromFile(product *Product, file os.File) (int, error) {
	catalog.lastId++
	product.id = catalog.lastId
	catalog.products[catalog.lastId] = product

	return catalog.lastId, nil
}

func (catalog *FileCatalog) DeleteElemetById(checkId int, Catalog *FileCatalog, reader bufio.Reader, file os.File) {

	_, _ = file.Seek(0, 0)
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

			_ = file.Truncate(int64(thisPosition)) // обрезаем файл по последнему байту этой строки

			_, _ = file.Seek(int64(thisPosition), 0)

			for i := checkId + 1; i <= Catalog.lastId; i++ { // в конец строки пушим из памяти все данные из виртуальной бд
				x := Catalog.products[i].name
				y := strconv.Itoa(int(Catalog.products[i].count))
				z := strconv.Itoa(int(Catalog.products[i].price))
				w := strconv.Itoa(Catalog.products[i].id - 1)

				str := "#" + w + "|" + x + "$" + y + "&" + z + "\n"

				_, _ = file.WriteString(str)
			}
		}
	}
}

func (catalog *FileCatalog) GetAll() map[int]*Product {
	// todo read all from file
	return catalog.products
}
*/

//It's OK
func SetCatalogType() Catalog {
	var catalog Catalog
	//var useFile string

	/*fmt.Print("Use file? ( y/n )")
	_, _ = fmt.Fscanln(os.Stdin, &useFile)

	if useFile == "y" {*/
	catalog = newFileCatalog()
	fmt.Println("localhost started with FileCatalog")
	return catalog
	/*}

	if useFile == "n" {
		catalog = NewInMemoryCatalog()
		fmt.Println("localhost started with InMemoryCatalog")
		return catalog
	} else {
		SetCatalogType()
	}*/
	return catalog
}

//It's OK
func openOrCreateFile() *os.File {

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

func createNewProduct(id int, name string, count int64, price float64) (*Product, error) {
	if name == "" || count < 0 || price < 0 {
		return nil, errors.New("invalid product data")
	} else {
		product := Product{id, name, count, price}

		return &product, nil
	}
}

func CheckError(r http.Request, name string, countErr, priceErr error) (hasError bool, form *CreateProductForm) {

	hasError = false

	createProductForm := CreateProductForm{}

	createProductForm.name = name
	createProductForm.count = r.FormValue("Second")
	createProductForm.price = r.FormValue("Third")

	if name == "" {
		createProductForm.nameError = "Ошибка имени"

		hasError = true
	}

	if countErr != nil {
		createProductForm.countError = "Ошибка колличества"
		createProductForm.count = r.FormValue("Second")

		hasError = true
	}

	if priceErr != nil {
		createProductForm.priceError = "Ошибка стоимости"
		createProductForm.price = r.FormValue("Third")

		hasError = true
	}

	return hasError, &createProductForm
}
