package Catalog

import (
	"Anton/View"
	"Anton/utils"
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	id    int
	name  string
	count int64
	price float64
}

type Catalog interface {
	AddNewProduct(*Product) (int, error)

	DeleteProductById(int) error

	GetAll() map[int]*Product

	EditProduct(*Product) (int, error)

	GetProductByID(int) (*Product, error)
}

//
//
//

type FilesCatalog struct {
	products map[int]*Product
	//file     os.File
	//lastId int
}

func NewFileCatalog() *FilesCatalog {

	catalog := FilesCatalog{}
	catalog.products = make(map[int]*Product)

	return &catalog
}

//
//
//

func (catalog FilesCatalog) AddNewProduct(product *Product) (int, error) {

	file := utils.OpenOrCreateFile()
	reader := bufio.NewReader(file)
	lastId := 1
	for {

		line, _, _ := reader.ReadLine()

		if len(line) == 0 {
			break
		}
		lastId++
	}
	a := strconv.Itoa(lastId)
	b := product.name
	c := strconv.Itoa(int(product.count))
	d := strconv.Itoa(int(product.price))

	textImpression := "#" + a + "|" + b + "$" + c + "&" + d + "\n"

	_, _ = file.Seek(0, 2) // устанавливаем курсор в позицию записи

	_, _ = file.WriteString(textImpression) //записываем введенную строку

	file.Close()

	return product.id, errors.New("cannot add product")
}

func (catalog FilesCatalog) DeleteProductById(cameId int) error {
	file := utils.OpenOrCreateFile()
	reader := bufio.NewReader(file)
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

		if cameId == id {

			buffer := ""

			for {
				nextLine, _, _ := reader.ReadLine()

				if len(nextLine) == 0 {
					break
				}

				i := strings.IndexAny(string(nextLine), "|")
				n := strings.IndexAny(string(nextLine), "$")
				c := strings.IndexAny(string(nextLine), "&")

				id, _ := strconv.Atoi(string(nextLine[1:i]))
				name := string(nextLine[i+1 : n])
				count, _ := strconv.ParseInt(string(nextLine[n+1:c]), 10, 64)
				price, _ := strconv.ParseFloat(string(nextLine[c+1:]), 64)

				x := name
				y := strconv.Itoa(int(count))
				z := strconv.Itoa(int(price))
				w := strconv.Itoa(id - 1)

				str := "#" + w + "|" + x + "$" + y + "&" + z + "\n"

				buffer += str
			}

			_, _ = file.Seek(int64(thisPosition), 0) // устанавливаем курсор в позицию записи + int64(len(textImpression))

			_ = file.Truncate(int64(thisPosition)) // обрезаем файл по последнему байту этой строки

			_, _ = file.WriteString(buffer)

		}

	}
	file.Close()
	return errors.New("can't edit product")
}

func (catalog FilesCatalog) GetAll() map[int]*Product {

	file := utils.OpenOrCreateFile()

	reader := bufio.NewReader(file)

	thisMap := catalog.products

	for id := range thisMap {
		delete(thisMap, id)
	}

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

		thisMap[id] = &Product{id, name, count, price}
	}

	file.Close()

	return thisMap
}

func (FilesCatalog) EditProduct(product *Product) (int, error) {
	file := utils.OpenOrCreateFile()
	reader := bufio.NewReader(file)
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

		if product.id == id {

			var buffer string

			for {
				nextLine, _, _ := reader.ReadLine()

				if len(nextLine) == 0 {
					break
				}

				buffer += string(nextLine) + "\n"
			}

			a := strconv.Itoa(product.id)
			b := product.name
			c := strconv.Itoa(int(product.count))
			d := strconv.Itoa(int(product.price))

			str := "#" + a + "|" + b + "$" + c + "&" + d + "\n"

			_, _ = file.Seek(int64(thisPosition), 0) // устанавливаем курсор в позицию записи + int64(len(textImpression))

			_, _ = file.WriteString(str) //записываем введенную строку

			_ = file.Truncate(int64(thisPosition + len(str))) // обрезаем файл по последнему байту этой строки

			_, _ = file.WriteString(buffer)

			file.Close()
		}
	}
	return product.id, errors.New("can't edit product")
}

func (FilesCatalog) GetProductByID(cameId int) (*Product, error) {

	file := utils.OpenOrCreateFile()
	reader := bufio.NewReader(file)
	_, _ = file.Seek(0, 0)

	for {
		line, _, _ := reader.ReadLine()

		if len(line) == 0 {
			break
		}

		id, _ := strconv.Atoi(string(line[1:strings.IndexAny(string(line), "|")]))

		if cameId == id {

			i := strings.IndexAny(string(line), "|")
			n := strings.IndexAny(string(line), "$")
			c := strings.IndexAny(string(line), "&")

			id, _ := strconv.Atoi(string(line[1:i]))
			name := string(line[i+1 : n])
			count, _ := strconv.ParseInt(string(line[n+1:c]), 10, 64)
			price, _ := strconv.ParseFloat(string(line[c+1:]), 64)

			return &Product{id, name, count, price}, nil
		}
	}
	return nil, errors.New("product not find")
}

//
//
//
//
//

type InMemoryCatalog struct {
	products map[int]*Product
	//lastID   int
}

func NewInMemoryCatalog() *InMemoryCatalog {

	catalog := InMemoryCatalog{}
	catalog.products = make(map[int]*Product)

	return &catalog
}

func (catalog InMemoryCatalog) AddNewProduct(product *Product) (int, error) {

	a := len(catalog.products)
	product.id = a + 1
	catalog.products[a+1] = product

	return 0, errors.New("cannot add product")
}

func (catalog InMemoryCatalog) DeleteProductById(cameId int) error {

	//

	for key := cameId; key < len(catalog.products); key++ {
		catalog.products[key] = catalog.products[key+1]
		catalog.products[key].id--

	}
	delete(catalog.products, len(catalog.products))

	return errors.New("can't edit product")
}

func (catalog InMemoryCatalog) GetAll() map[int]*Product {
	return catalog.products
}

func (catalog InMemoryCatalog) EditProduct(product *Product) (int, error) {

	catalog.products[product.id] = product

	return product.id, errors.New("can't edit product")
}

func (catalog InMemoryCatalog) GetProductByID(cameId int) (*Product, error) {
	return catalog.products[cameId], errors.New("product not found")
}

func SetCatalogType() Catalog {

	var catalog Catalog

	useFile := os.Args[1]

	if useFile == "f" {
		catalog = NewFileCatalog()
		fmt.Println("localhost started with FileCatalog")
		return catalog
	}

	if useFile == "m" {
		catalog = NewInMemoryCatalog()
		fmt.Println("localhost started with InMemoryCatalog")
		return catalog
	} else {
		SetCatalogType()
	}
	return catalog
}

//It's OK
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

func CreateNewProduct(id int, name string, count int64, price float64) (*Product, error) {
	if name == "" || count < 0 || price < 0 {
		return nil, errors.New("invalid product data")
	} else {
		product := Product{id, name, count, price}

		return &product, nil
	}
}

func CheckError(r http.Request, name string, countErr, priceErr error) (hasError bool, form *View.CreateProductForm) {

	hasError = false

	createProductForm := View.CreateProductForm{}

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
