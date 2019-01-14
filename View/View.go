package View

import (
	"AmazingCatalog/CatalogModel"
	"fmt"
	"strconv"
	"strings"
)

var styles = `
<style>
	html{
		display: flex; 
		justify-content: center;
		}
	caption{
		font-weight: bold;
		margin: 20px 0px 30px 0px;
		}
		
		table {
			border-collapse: collapse;
  	  }
    	td {
     	  border: 1px solid black; 
  	 }
		button{
		border: 0;
		width: 100%;
		color: white;
		background:deeppink;
		}
	form{
		display: flex;
		flex-direction: column;
	}
	
</style>`

type CreateProductForm struct {
	Name      string
	NameError string

	Count      string
	CountError string

	Price      string
	PriceError string
}

func renderError(err string) string {
	if err == "" {
		return ""
	}
	return "<span style='color:red'>" + err + "</span>"
}

func html(child ...string) string {
	return fmt.Sprint(`
	<!DOCTYPE html>
		<html>`,
		childIterator(child),
		`</html>`)
}

func head(child ...string) string {
	return fmt.Sprint(`
	<head>
    	<meta charset="utf-8">`,
		childIterator(child),
		`</head>`)
}

func body(child ...string) string {
	return fmt.Sprint(`
	<body>`,
		childIterator(child),
		`</body>`)
}

func title(child ...string) string {
	return fmt.Sprint(`
	<title>`,
		childIterator(child),
		`</title>`)
}

func childIterator(str []string) string {
	newStr := ""
	for i := range str {
		newStr += str[i]
	}

	return newStr
}

func PrintProductList(catalog CatalogModel.Catalog) string {

	b := ""
	mep := catalog.GetAll()

	for i := range mep {

		arr := make([]string, 13)
		arr[0] = `<tr><td>`
		arr[1] = strconv.Itoa(mep[i].GetId())
		arr[2] = `</td><td>`
		arr[3] = mep[i].GetName()
		arr[4] = `</td><td>`
		arr[5] = strconv.FormatInt(mep[i].GetCount(), 10)
		arr[6] = `</td><td>`
		arr[7] = strconv.FormatFloat(mep[i].GetPrice(), 'f', 0, 64)
		arr[8] = `</td><td>`
		arr[9] = "" //strconv.FormatFloat(mep[i].GetViewPrice(), 'f', 0, 64)
		arr[10] = `</td><td><a href="http://localhost:8080/edit?product_id=` + strconv.Itoa(i) + `"><button>Изменить</button></a></td>`
		arr[11] = `<td><a href="http://localhost:8080/delete?product_id=` + strconv.Itoa(i) + `"><button>Удалить</button></a></td>`
		arr[12] = `</tr>`

		b += strings.Join(arr, "")
	}

	return fmt.Sprint(
		html(
			head(
				title(`Окно результатов`),
				styles),
			body(`<div>
			<table>
			<caption>Список товаров</caption>
				<tr>
					<td>Id</td>
					<td>Наименование</td>
					<td>Колличество</td>
					<td>Цена</td>
					<td>Стоимость с учетом скидки</td>
					<td>Редактировать</td>
					<td>Удалить</td>
				</tr>
				`, b, `
			</table>
				<br>
			<form action="http://localhost:8080/list" method="POST">
				<a href="http://localhost:8080/list"><button>Добавить</button> </a>
			</form>`)))
}

// todo pass by ref
func AddPageView(form CreateProductForm, headerName, buttonName string) string {
	return fmt.Sprint(
		html(
			head(
				title("Добавление продукта"),
				styles),
			body(`
			<h3>`, headerName, ` </h3>
			<form action="http://localhost:8080/add" method="POST">
				 <br>
				<input type="text" Name="First" placeholder="Наименование" value="`, form.Name, `">`, renderError(form.NameError), `
				 <br>
				<input type="text" Name="Second" placeholder="Колличество" value="`, form.Count, `">`, renderError(form.PriceError), `
				 <br>
				<input type="text" Name="Third" placeholder="Стоимость" value="`, form.Price, `">`, renderError(form.PriceError), `
			 	 <br>
				<select Name ="productType" value="productType">
					<option> Фрукты </option>
					<option> Овощи </option>
					<option> Мясо </option>
				</select> 
				 <br>
				<a href="http://localhost:8080/list"><button Name = "product_id" >`, buttonName, `</button></a>
			</form>`)))
}

func EditPageView(form CreateProductForm, buttonName string, id int) string {
	return fmt.Sprint(
		html(
			head(
				title(`Редактирвание продукта`),
				styles),
			body(`
			<h3>Измените продукт</h3>
			<form action="http://localhost:8080/edit?product_id=`+strconv.Itoa(id)+` " method="POST">
		    	 <br>
				<input type="text" Name="First" placeholder="Наименование" value="`, form.Name, `">`, renderError(form.NameError), `
				 <br>
				<input type="text" Name="Second" placeholder="Колличество" value="`, form.Count, `">`, renderError(form.CountError), `
				 <br>
				<input type="text" Name="Third" placeholder="Стоимость" value="`, form.Price, `">`, renderError(form.PriceError), `
			 	<br>
				<button Name = "product_id" value ="`+strconv.Itoa(id)+`">`, buttonName, `</button>
			</form>`)))
}
