package CatalogModel

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FilesCatalog struct {
	products map[int]*Product
}

type FileCatalogFactory struct {
}

func NewFileCatalogFactory() FileCatalogFactory {
	FileCatalogFactory := FileCatalogFactory{}

	return FileCatalogFactory
}

func (FileCatalogFactory) CreateCatalog() Catalog {
	catalog := FilesCatalog{}
	catalog.products = make(map[int]*Product)

	return &catalog
}

//

func (catalog *FilesCatalog) AddNewProduct(product *Product) (int, error) {

	//todo move to method withName getLastId lastId := catalog.getLastId()
	file, reader := OpenOrCreateFile()

	lastId := 1
	for {
		line, _, _ := reader.ReadLine()
		if len(line) == 0 {
			break
		}
		lastId++
	}

	str := "#" + strconv.Itoa(lastId) + "|" + product.name + "$" + strconv.Itoa(int(product.count)) + "&" + strconv.Itoa(int(product.price)) + ":" + product.productType + "\n"

	_, _ = file.Seek(0, 2) // устанавливаем курсор в позицию записи

	_, _ = file.WriteString(str) //записываем введенную строку

	_ = file.Close()

	return product.id, errors.New("cannot add product")
}

func (catalog *FilesCatalog) DeleteProductById(cameId int) error {

	file, reader := OpenOrCreateFile()

	currentDisplacement := 0

	for {

		nextLine, _, _ := reader.ReadLine()
		lineX := string(nextLine) + "\n"
		StartCurrentString := len(lineX)
		currentDisplacement += len(lineX)
		thisPosition := currentDisplacement - StartCurrentString
		if len(nextLine) == 0 {
			break
		}

		id, _ := strconv.Atoi(string(nextLine[1:strings.IndexAny(string(nextLine), "|")]))

		product := &Product{}

		if cameId == id {

			buffer := ""

			for {
				line, _, _ := reader.ReadLine()

				if len(line) == 0 {
					break
				}

				product = CreateProductFromFile(line, product)

				str := "#" + strconv.Itoa(product.id-1) + "|" + product.name + "$" + strconv.Itoa(int(product.count)) + "&" + strconv.Itoa(int(product.price)) + ":" + product.productType + "\n"
				buffer += str
			}

			_, _ = file.Seek(int64(thisPosition), 0) // устанавливаем курсор в позицию записи + int64(len(textImpression))

			_ = file.Truncate(int64(thisPosition)) // обрезаем файл по последнему байту этой строки

			_, _ = file.WriteString(buffer)

		}

	}
	_ = file.Close()

	return errors.New("can't edit product")
}

func (catalog *FilesCatalog) GetAll() map[int]*Product {

	file, reader := OpenOrCreateFile()

	for id := range catalog.products {
		delete(catalog.products, id)
	}
	product := &Product{}

	for {

		line, _, _ := reader.ReadLine()

		if len(line) == 0 {
			break
		}
		product := CreateProductFromFile(line, product)
		catalog.products[product.id] = product
	}

	_ = file.Close()

	return catalog.products
}

func (*FilesCatalog) EditProduct(cameId int, name string, count int64, price float64) (int, error) {

	file, reader := OpenOrCreateFile()

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

				buffer += string(nextLine) + "\n"
			}
			fmt.Println(string(line))

			productType := line[strings.IndexAny(string(line), ":"):]

			str := "#" + strconv.Itoa(cameId) + "|" + name + "$" + strconv.Itoa(int(count)) + "&" + strconv.Itoa(int(price)) + string(productType) + "\n"

			_, _ = file.Seek(int64(thisPosition), 0) // устанавливаем курсор в позицию записи + int64(len(textImpression))

			_, _ = file.WriteString(str) //записываем введенную строку

			_ = file.Truncate(int64(thisPosition + len(str))) // обрезаем файл по последнему байту этой строки

			_, _ = file.WriteString(buffer)

			_ = file.Close()
		}
	}
	return cameId, errors.New("can't edit product")
}

func (*FilesCatalog) GetProductByID(cameId int) (*Product, error) {

	file, reader := OpenOrCreateFile()

	product := &Product{}

	for {
		line, _, _ := reader.ReadLine()

		if len(line) == 0 {
			break
		}

		id, _ := strconv.Atoi(string(line[1:strings.IndexAny(string(line), "|")]))

		if cameId == id {

			product = CreateProductFromFile(line, product)
		}
	}
	_ = file.Close()
	return product, errors.New("product not found")
}

// open or create file and set read position at 0
func OpenOrCreateFile() (*os.File, *bufio.Reader) {
	file := &os.File{}

	_, err := os.Stat("MyFile.txt")
	if err != nil {
		file, _ = os.Create("MyFile.txt")
		fmt.Println("I create File")
	} else {
		file, _ = os.OpenFile("MyFile.txt", os.O_RDWR, 111)
		fmt.Println("I open File")
	}
	reader := bufio.NewReader(file)
	_, _ = file.Seek(0, 0)
	return file, reader
}

func CreateProductFromFile(line []byte, product *Product) *Product {
	a := strings.IndexAny(string(line), "|")
	b := strings.IndexAny(string(line), "$")
	c := strings.IndexAny(string(line), "&")
	d := strings.IndexAny(string(line), ":")

	id, _ := strconv.Atoi(string(line[1:a]))
	name := string(line[a+1 : b])
	count, _ := strconv.ParseInt(string(line[b+1:c]), 10, 64)
	price, _ := strconv.ParseFloat(string(line[c+1:d]), 64)
	productType := string(line[d+1:])

	product = &Product{id, name, productType, count, price}
	return product
}
