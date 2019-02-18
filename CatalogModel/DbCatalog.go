package CatalogModel

import (
	"database/sql"
	"github.com/pkg/errors"
)

type DbCatalog struct {
}

type DBCatalogFactory struct {
}

func NewDBCatalogFactory() DBCatalogFactory {
	DBCatalogFactory := DBCatalogFactory{}
	return DBCatalogFactory
}

func (DBCatalogFactory) CreateCatalog() Catalog {
	catalog := DbCatalog{}
	return &catalog
}

func (catalog *DbCatalog) AddNewProduct(product *Product) (int, error) {
	db, err := sql.Open("postgres", "user = postgres password = 123 dbname = Catalog sslmode = disable")
	id := 0
	row := db.QueryRow("select max(id) from catalog")
	err = row.Scan(&id)

	if id == 0 {
		id = 1
	} else {
		id = int(id) + 1
	}

	_, err = db.Exec("insert into catalog (id, name, count, price, producttype) values ($1,$2,$3,$4,$5)",
		id, product.name, product.count, product.price, product.productType)
	err = db.Close()
	return product.id, errors.Wrap(err, "AddNewProduct")
}

func (catalog *DbCatalog) DeleteProductById(cameId int) error {
	db, err := sql.Open("postgres", "user = postgres password = 123 dbname = Catalog sslmode = disable")
	maxId := 0
	row := db.QueryRow("select max(id) from catalog")
	err = row.Scan(&maxId)
	_, err = db.Exec("delete from catalog where id = $1", cameId) //удаляем продукт по id

	for i := cameId + 1; i <= maxId; i++ {
		_, err = db.Exec("update catalog set id = $1 where id = $2", i-1, i)
	}
	err = db.Close()
	return errors.Wrap(err, "DeleteProductById")
}

func (catalog *DbCatalog) GetAll() (map[int]*Product, error) {

	db, err := sql.Open("postgres", "user = postgres password = 123 dbname = Catalog sslmode = disable")
	rows, err := db.Query("select * from catalog")
	thisMap := map[int]*Product{}
	i := 1
	for rows.Next() {
		product := Product{}
		err = rows.Scan(&product.id, &product.name, &product.count, &product.price, &product.productType)
		thisMap[i] = &product
		i++
	}

	return thisMap, errors.Wrap(err, "GetAll")
}

func (*DbCatalog) EditProduct(cameId int, name string, count int64, price float64) (int, error) {
	db, err := sql.Open("postgres", "user = postgres password = 123 dbname = Catalog sslmode = disable")

	err = db.Ping()
	_ = db.QueryRow("update  catalog  set name = $1,count= $2,price= $3 where id = $4", name, count, price, cameId)
	return cameId, errors.Wrap(err, "DbCatalogEditProduct")
}

func (*DbCatalog) GetProductByID(cameId int) (*Product, error) {
	product := &Product{}

	db, err := sql.Open("postgres", "user = postgres password = 123 dbname = Catalog sslmode = disable")

	row := db.QueryRow("select id,name,count,price,producttype from catalog where id = $1", cameId)

	err = row.Scan(&product.id, &product.name, &product.count, &product.price, &product.productType)
	err = db.Close()

	return product, errors.Wrap(err, "GetProductByID")
}
