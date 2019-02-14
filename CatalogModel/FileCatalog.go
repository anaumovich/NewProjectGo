package CatalogModel

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
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
		line, _, err := reader.ReadLine()
		if len(line) == 0 {
			break
		}
		lastId++
		if err != nil {
			panic(errors.Wrap(err, "AddNewProduct:readLine"))
		}
	}

	str := "#" + strconv.Itoa(lastId) + "|" + product.name + "$" + strconv.Itoa(int(product.count)) +
		"&" + strconv.Itoa(int(product.price)) + ":" + product.productType + "\n"

	_, err := file.Seek(0, 2) // устанавливаем курсор в позицию записи

	_, err = file.WriteString(str) //записываем введенную строку

	err = file.Close()

	return product.id, errors.Wrap(err, "AddNewProduct")
}

func (catalog *FilesCatalog) DeleteProductById(cameId int) error {

	file, reader := OpenOrCreateFile()

	currentDisplacement := 0

	for {

		nextLine, _, err := reader.ReadLine()
		lineX := string(nextLine) + "\n"
		StartCurrentString := len(lineX)
		currentDisplacement += len(lineX)
		thisPosition := currentDisplacement - StartCurrentString
		if len(nextLine) == 0 {
			break
		}

		id, err := strconv.Atoi(string(nextLine[1:strings.IndexAny(string(nextLine), "|")]))

		product := &Product{}

		if cameId == id {

			buffer := ""

			for {
				line, _, err := reader.ReadLine()

				if len(line) == 0 {
					break
				}

				product = CreateProductFromFile(line, product)

				str := "#" + strconv.Itoa(product.id-1) + "|" + product.name + "$" + strconv.Itoa(int(product.count)) +
					"&" + strconv.Itoa(int(product.price)) + ":" + product.productType + "\n"
				buffer += str
				if err != nil {
					panic(errors.Wrap(err, "DeleteProductById"))
				}

			}

			_, err = file.Seek(int64(thisPosition), 0) // устанавливаем курсор в позицию записи + int64(len(textImpression))

			err = file.Truncate(int64(thisPosition)) // обрезаем файл по последнему байту этой строки

			_, err = file.WriteString(buffer)
			if err != nil {
				panic(errors.Wrap(err, "DeleteProductById"))
			}
		}

	}
	err := file.Close()

	return errors.Wrap(err, "DeleteProductById")
}

func (catalog *FilesCatalog) GetAll() (map[int]*Product, error) {

	file, reader := OpenOrCreateFile()

	for id := range catalog.products {
		delete(catalog.products, id)
	}
	product := &Product{}

	for {

		line, _, err := reader.ReadLine()

		if len(line) == 0 {
			break
		}
		product := CreateProductFromFile(line, product)
		catalog.products[product.id] = product
		if err != nil {
			panic(errors.Wrap(err, "GetAll"))
		}
	}

	err := file.Close()

	return catalog.products, errors.Wrap(err, "GetAll")
}

func (*FilesCatalog) EditProduct(cameId int, name string, count int64, price float64) (int, error) {

	file, reader := OpenOrCreateFile()

	currentDisplacement := 0

	for {

		line, _, err := reader.ReadLine()
		lineX := string(line) + "\n"
		StartCurrentString := len(lineX)
		currentDisplacement += len(lineX)
		thisPosition := currentDisplacement - StartCurrentString

		if len(line) == 0 {
			break
		}

		id, err := strconv.Atoi(string(line[1:strings.IndexAny(string(line), "|")]))

		if cameId == id {

			buffer := ""

			for {
				nextLine, _, err := reader.ReadLine()

				if len(nextLine) == 0 {
					break

				}

				buffer += string(nextLine) + "\n"
				if err != nil {
					panic(errors.Wrap(err, "EditProduct"))
				}
			}
			fmt.Println(string(line))

			productType := line[strings.IndexAny(string(line), ":"):]

			str := "#" + strconv.Itoa(cameId) + "|" + name + "$" + strconv.Itoa(int(count)) +
				"&" + strconv.Itoa(int(price)) + string(productType) + "\n"

			_, err = file.Seek(int64(thisPosition), 0) // устанавливаем курсор в позицию записи + int64(len(textImpression))

			_, err = file.WriteString(str) //записываем введенную строку

			err = file.Truncate(int64(thisPosition + len(str))) // обрезаем файл по последнему байту этой строки

			_, err = file.WriteString(buffer)

			if err != nil {
				panic(errors.Wrap(err, "EditProduct"))
			}
		}
	}
	err := file.Close()
	return cameId, errors.Wrap(err, "EditProduct")
}

func (*FilesCatalog) GetProductByID(cameId int) (*Product, error) {

	file, reader := OpenOrCreateFile()

	product := &Product{}

	for {
		line, _, err := reader.ReadLine()

		if len(line) == 0 {
			break
		}

		id, err := strconv.Atoi(string(line[1:strings.IndexAny(string(line), "|")]))

		if cameId == id {

			product = CreateProductFromFile(line, product)
		}
		if err != nil {
			panic(errors.Wrap(err, "GetProductByID"))
		}
	}
	err := file.Close()
	return product, errors.Wrap(err, "GetProductByID")
}

// open or create file and set read position at 0
func OpenOrCreateFile() (*os.File, *bufio.Reader) {
	file := &os.File{}

	_, err := os.Stat("MyFile.txt")
	if err != nil {
		file, err = os.Create("MyFile.txt")
		if err != nil {
			panic(errors.Wrap(err, "OpenOrCreateFile"))
		}
		fmt.Println("I create File")
	} else {
		file, err = os.OpenFile("MyFile.txt", os.O_RDWR, 111)
		if err != nil {
			panic(errors.Wrap(err, "OpenOrCreateFile"))
		}
		fmt.Println("I open File")
	}
	reader := bufio.NewReader(file)
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(errors.Wrap(err, "OpenOrCreateFile"))
	}
	return file, reader
}

func CreateProductFromFile(line []byte, product *Product) *Product {
	a := strings.IndexAny(string(line), "|")
	b := strings.IndexAny(string(line), "$")
	c := strings.IndexAny(string(line), "&")
	d := strings.IndexAny(string(line), ":")

	id, err := strconv.Atoi(string(line[1:a]))
	name := string(line[a+1 : b])
	count, err := strconv.ParseInt(string(line[b+1:c]), 10, 64)
	price, err := strconv.ParseFloat(string(line[c+1:d]), 64)
	productType := string(line[d+1:])
	if err != nil {
		panic(errors.Wrap(err, "CreateProductFromFile"))
	}
	product = &Product{id, name, productType, count, price}
	return product
}
