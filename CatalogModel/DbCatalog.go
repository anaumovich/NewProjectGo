package CatalogModel

import (
	"database/sql"
	"errors"
)

type DbCatalog struct{}

func NewDbCatalog() *DbCatalog {

	catalog := DbCatalog{}
	return &catalog
}

func (catalog DbCatalog) AddNewProduct(product *product) (int, error) {
	db, _ := sql.Open("postgres", "user = postgres password = 123 dbname = Catalog sslmode = disable")
	id := 0
	row := db.QueryRow("select max(id) from catalog")
	_ = row.Scan(&id)

	if id == 0 {
		id = 1
	} else {
		id = int(id) + 1
	}

	_, _ = db.Exec("insert into catalog (id, name, count, price, producttype) values ($1,$2,$3,$4,$5)",
		id, product.name, product.count, product.price, product.productType)
	_ = db.Close()
	return product.id, errors.New("cannot add product")
}

func (catalog DbCatalog) DeleteProductById(cameId int) error {
	db, _ := sql.Open("postgres", "user = postgres password = 123 dbname = Catalog sslmode = disable")
	maxId := 0
	row := db.QueryRow("select max(id) from catalog")
	_ = row.Scan(&maxId)
	_, _ = db.Exec("delete from catalog where id = $1", cameId) //удаляем продукт по id

	for i := cameId + 1; i <= maxId; i++ {
		_, _ = db.Exec("update catalog set id = $1 where id = $2", i-1, i)
	}
	_ = db.Close()
	return errors.New("can't edit product")
}

func (catalog DbCatalog) GetAll() map[int]*product {

	db, _ := sql.Open("postgres", "user = postgres password = 123 dbname = Catalog sslmode = disable")
	rows, _ := db.Query("select * from catalog")
	thisMap := map[int]*product{}
	i := 1
	for rows.Next() {
		product := product{}
		_ = rows.Scan(&product.id, &product.name, &product.count, &product.price, &product.productType)
		thisMap[i] = &product
		i++
	}
	_ = db.Close()
	return thisMap
}

func (DbCatalog) EditProduct(cameId int, name string, count int64, price float64) (int, error) {
	db, _ := sql.Open("postgres", "user = postgres password = 123 dbname = Catalog sslmode = disable")
	_ = db.QueryRow("update  catalog  set name = $1,count= $2,price= $3 where id = $4", name, count, price, cameId)
	return cameId, errors.New("can't edit product")
}

func (DbCatalog) GetProductByID(cameId int) (*product, error) {
	Product := &product{}

	db, _ := sql.Open("postgres", "user = postgres password = 123 dbname = Catalog sslmode = disable")

	row := db.QueryRow("select id,name,count,price,producttype from catalog where id = $1", cameId)

	_ = row.Scan(&Product.id, &Product.name, &Product.count, &Product.price, &Product.productType)
	_ = db.Close()
	return Product, errors.New("product not found")
}
