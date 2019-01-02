package CatalogModel

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type product struct {
	id    int
	name  string
	count int64
	price float64
	CalculateDiscount CalculateDiscount
}

type Catalog interface {
	CreateNewProduct(int, string, int64, float64) (*product, error)

	AddNewProduct(*product) (int, error)

	DeleteProductById(int) error

	GetAll() map[int]*product

	EditProduct(int, string, int64, float64) (int, error)

	GetProductByID(int) (*product, error)
}

//

type FilesCatalog struct {
	products map[int]*product
}

func NewFileCatalog() *FilesCatalog {

	catalog := FilesCatalog{}
	catalog.products = make(map[int]*product)

	return &catalog
}

//

func (catalog FilesCatalog) CreateNewProduct(id int, name string, count int64, price float64) (*product, error) {
	if name == "" || count < 0 || price < 0 {
		return nil, errors.New("invalid product data")
	} else {

		MeatProduct:=

		product :=

		product.id = id
		product.name = name
		product.count = count
		product.price = price

		return &product, nil
	}
}

func (catalog FilesCatalog) AddNewProduct(product *product) (int, error) {

	file := OpenOrCreateFile()
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
	file := OpenOrCreateFile()
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

				str := "#" + strconv.Itoa(id-1) + "|" + name + "$" + strconv.Itoa(int(count)) + "&" + strconv.Itoa(int(price)) + "\n"

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

func (catalog FilesCatalog) GetAll() map[int]*product {

	file := OpenOrCreateFile()

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

		thisMap[id] = &product{id, name, count, price}

	}

	file.Close()

	return thisMap
}

func (FilesCatalog) EditProduct(cameId int, name string, count int64, price float64) (int, error) {
	file := OpenOrCreateFile()
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

			var buffer string

			for {
				nextLine, _, _ := reader.ReadLine()

				if len(nextLine) == 0 {
					break
				}

				buffer += string(nextLine) + "\n"
			}

			str := "#" + strconv.Itoa(cameId) + "|" + name + "$" + strconv.Itoa(int(count)) + "&" + strconv.Itoa(int(price)) + "\n"

			_, _ = file.Seek(int64(thisPosition), 0) // устанавливаем курсор в позицию записи + int64(len(textImpression))

			_, _ = file.WriteString(str) //записываем введенную строку

			_ = file.Truncate(int64(thisPosition + len(str))) // обрезаем файл по последнему байту этой строки

			_, _ = file.WriteString(buffer)

			file.Close()
		}
	}
	return cameId, errors.New("can't edit product")
}

func (FilesCatalog) GetProductByID(cameId int) (*product, error) {

	file := OpenOrCreateFile()
	reader := bufio.NewReader(file)
	_, _ = file.Seek(0, 0)

	Product := &product{}

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

			Product = &product{id, name, count, price}

		}
	}
	return Product, errors.New("product not found")
}

//

type InMemoryCatalog struct {
	products map[int]*product
}

func NewInMemoryCatalog() *InMemoryCatalog {

	catalog := InMemoryCatalog{}
	catalog.products = make(map[int]*product)

	return &catalog
}

func (catalog InMemoryCatalog) CreateNewProduct(id int, name string, count int64, price float64) (*product, error) {
	if name == "" || count < 0 || price < 0 {
		return nil, errors.New("invalid product data")
	} else {
		product := product{id, name, count, price}

		return &product, nil
	}
}

func (catalog InMemoryCatalog) AddNewProduct(product *product) (int, error) {

	a := len(catalog.products)
	product.id = a + 1
	catalog.products[a+1] = product

	return 0, errors.New("cannot add product")
}

func (catalog InMemoryCatalog) DeleteProductById(cameId int) error {
	for key := cameId; key < len(catalog.products); key++ {
		catalog.products[key] = catalog.products[key+1]
		catalog.products[key].id--
	}
	delete(catalog.products, len(catalog.products))

	return errors.New("can't edit product")
}

func (catalog InMemoryCatalog) GetAll() map[int]*product {
	return catalog.products
}

func (catalog InMemoryCatalog) EditProduct(cameId int, name string, count int64, price float64) (int, error) {

	catalog.products[cameId].id = cameId
	catalog.products[cameId].name = name
	catalog.products[cameId].count = count
	catalog.products[cameId].price = price

	return cameId, errors.New("can't edit product")
}

func (catalog InMemoryCatalog) GetProductByID(cameId int) (*product, error) {
	return catalog.products[cameId], errors.New("product not found")
}

//

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

func (p product) GetId() int {
	id := p.id
	return id
}
func (p product) GetName() string {
	name := p.name
	return name
}
func (p product) GetCount() int64 {
	count := p.count
	return count
}
func (p product) GetPrice() float64 {
	price := p.price
	return price
}
