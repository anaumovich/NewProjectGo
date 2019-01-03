package Strategy

/*
import "fmt"

type Catalogger interface {
	AddNewProduct(*Product) (int, error)
}

//структура имеет поле с именем productAdder  типа ProductAdder
type Catalog1 struct {
	productAdder ProductAdder
}

func (catalog *Catalog1) AddNewProduct(product *Product) (int, error) {

	catalog.productAdder.AddProduct(product)

	panic("implement me")
}

//функция конструктор
func NewsCatalog(add ProductAdder) *Catalog1 {
	return &Catalog1{productAdder: add}
}

type ProductAdder interface {
	AddProduct(*Product)
}

// модель поведения реализует интерфейс для добавления продукта в память
type CanAddProductInMemory struct {
}

func (CanAddProductInMemory) AddProduct(product *Product) {
	fmt.Println(Product{})
}

// модель поведения реализует интерфейс для добавления продукта в файл
type CanAddProductInFile struct {
}

func (CanAddProductInFile) AddProduct(product *Product) {
	fmt.Println(Product{})
}

func main() {

	InMemoryCatalog := NewsCatalog(CanAddProductInMemory{})

	FileCatalog := NewsCatalog(CanAddProductInFile{})

	InMemoryCatalog.productAdder.AddProduct(&Product{})
	FileCatalog.productAdder.AddProduct(&Product{})
}
*/
