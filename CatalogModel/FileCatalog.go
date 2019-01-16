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

	// todo serialize
	a := strconv.Itoa(lastId)
	b := product.name
	c := strconv.Itoa(int(product.count))
	d := strconv.Itoa(int(product.price))
	e := product.productType

	textImpression := "#" + a + "|" + b + "$" + c + "&" + d + ":" + e + "\n"

	_, _ = file.Seek(0, 2) // устанавливаем курсор в позицию записи

	// write to file
	_, _ = file.WriteString(textImpression) //записываем введенную строку

	_ = file.Close()

	return product.id, errors.New("cannot add product")
}

func (catalog *FilesCatalog) DeleteProductById(cameId int) error {
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

				a := strings.IndexAny(string(nextLine), "|")
				b := strings.IndexAny(string(nextLine), "$")
				c := strings.IndexAny(string(nextLine), "&")
				d := strings.IndexAny(string(nextLine), ":")

				id, _ := strconv.Atoi(string(nextLine[1:a]))
				name := string(nextLine[a+1 : b])
				count, _ := strconv.ParseInt(string(nextLine[b+1:c]), 10, 64)
				price, _ := strconv.ParseFloat(string(nextLine[c+1:d]), 64)
				productType := string(nextLine[d+1:])

				str := "#" + strconv.Itoa(id-1) + "|" + name + "$" + strconv.Itoa(int(count)) + "&" + strconv.Itoa(int(price)) + ":" + productType + "\n"

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

	file := OpenOrCreateFile()

	reader := bufio.NewReader(file)

	for id := range catalog.products {
		delete(catalog.products, id)
	}

	for {

		line, _, _ := reader.ReadLine()

		if len(line) == 0 {
			break
		}

		a := strings.IndexAny(string(line), "|")
		b := strings.IndexAny(string(line), "$")
		c := strings.IndexAny(string(line), "&")
		d := strings.IndexAny(string(line), ":")

		id, _ := strconv.Atoi(string(line[1:a]))
		name := string(line[a+1 : b])
		count, _ := strconv.ParseInt(string(line[b+1:c]), 10, 64)
		price, _ := strconv.ParseFloat(string(line[c+1:d]), 64)
		productType := string(line[d+1:])

		catalog.products[id] = &Product{id, name, productType, count, price}

	}

	_ = file.Close()

	return catalog.products
}

func (*FilesCatalog) EditProduct(cameId int, name string, count int64, price float64) (int, error) {
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

			b := "#" + strconv.Itoa(cameId) + "|" + name + "$" + strconv.Itoa(int(count)) + "&" + strconv.Itoa(int(price))

			productType := line[len(b)-1:]

			fmt.Println(string(productType))

			str := "#" + strconv.Itoa(cameId) + "|" + name + "$" + strconv.Itoa(int(count)) + "&" + strconv.Itoa(int(price)) + ":" + string(productType) + "\n"

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

	file := OpenOrCreateFile()
	reader := bufio.NewReader(file)
	_, _ = file.Seek(0, 0)

	product := &Product{}

	for {
		line, _, _ := reader.ReadLine()

		if len(line) == 0 {
			break
		}

		id, _ := strconv.Atoi(string(line[1:strings.IndexAny(string(line), "|")]))

		if cameId == id {

			a := strings.IndexAny(string(line), "|")
			b := strings.IndexAny(string(line), "$")
			c := strings.IndexAny(string(line), "&")
			d := strings.IndexAny(string(line), ":")

			id, _ := strconv.Atoi(string(line[1:a]))
			name := string(line[a+1 : b])
			count, _ := strconv.ParseInt(string(line[b+1:c]), 10, 64)
			price, _ := strconv.ParseFloat(string(line[c+1:d]), 64)
			productType := string(line[d+1:])

			product.id = id
			product.name = name
			product.count = count
			product.price = price
			product.productType = productType

		}
	}
	return product, errors.New("product not found")
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
