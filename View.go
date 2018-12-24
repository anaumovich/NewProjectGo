package main

import (
	"fmt"
	"strconv"
	"strings"
)

type CreateProductForm struct {
	name      string
	nameError string

	count      string
	countError string

	price      string
	priceError string
}

func renderError(err string) string {
	if err == "" {
		return ""
	}
	return "<span style='color:red'>" + err + "</span>"
}

func generatePageListHTMLController(MyCatalog *Catalog) string {

	b := ""

	for a := 1; a <= MyCatalog.lastId; a++ {
		arr := make([]string, 10)
		arr[0] = `<tr><td>`
		arr[1] = strconv.Itoa(MyCatalog.products[a].id)
		arr[2] = `</td><td>`
		arr[3] = MyCatalog.products[a].name
		arr[4] = `</td><td>`
		arr[5] = strconv.FormatInt(MyCatalog.products[a].count, 10)
		arr[6] = `</td><td>`
		arr[7] = strconv.FormatFloat(MyCatalog.products[a].price, 'f', 0, 64)
		arr[8] = `</td><td><a href="http://localhost:8080/edit?product_id=` + strconv.Itoa(a) + `"><button>Изменить</button></a></td>`
		arr[9] = `</td></tr>`
		b += strings.Join(arr, "")
	}

	return b
}

// todo pass by ref
func AddPageView(form CreateProductForm, headerName, buttonName string, id int, localWay string) string {
	return fmt.Sprint(`
<!DOCTYPE html>
<html> 
<style>
	html{
		display: flex; 
		justify-content: center;
		}
	button{
		border: 0;
		width: 100%;
		color: white;
		background:deeppink;
		}
	
</style>
<head>
    <meta charset="utf-8">
    <title>Кнопка отправки формы</title>
    </head>
<body>
<h3>`, headerName, ` </h3>
<form action="http://localhost:8080/add`, localWay, ` " method="POST" style="display: flex; flex-direction: column;">
	<br>
    <input type="text" name="First" placeholder="Наименование" value="`, form.name, `">`, renderError(form.nameError), `
	<br>
    <input type="text" name="Second" placeholder="Колличество" value="`, form.count, `">`, renderError(form.countError), `
    <br>
	<input type="text" name="Third" placeholder="Колличество" value="`, form.price, `">`, renderError(form.priceError), `
	<br>
	<a href="http://localhost:8080/list"><button name = "product_id" value ="`+strconv.Itoa(id)+`">`, buttonName, `</button></a>

</form>
</body>
</html>
`)
}

var styles = `
<style>
	html{
		display: flex; 
		justify-content: center;
		}
	button{
		border: 0;
		width: 100%;
		color: white;
		background:deeppink;
		}
	
</style>`

func EditPageView(form CreateProductForm, headerName, buttonName string, id int) string {
	return fmt.Sprint(`
<!DOCTYPE html>
<html> 
`, styles, `
<head>
    <meta charset="utf-8">
    <title>Кнопка отправки формы</title>
    </head>
<body>
<h3>`, headerName, ` </h3>
<form action="http://localhost:8080/edit?product_id=`+strconv.Itoa(id)+` " method="POST" style="display: flex; flex-direction: column;">
	<br>
    <input type="text" name="First" placeholder="Наименование" value="`, form.name, `">`, renderError(form.nameError), `
	<br>
    <input type="text" name="Second" placeholder="Колличество" value="`, form.count, `">`, renderError(form.countError), `
    <br>
	<input type="text" name="Third" placeholder="Колличество" value="`, form.price, `">`, renderError(form.priceError), `
	<br>
	<a href="http://localhost:8080/edit?product_id=`+strconv.Itoa(id)+`"><button name = "product_id" value ="`+strconv.Itoa(id)+`">`, buttonName, `</button></a>

</form>
</body>
</html>
`)
}

func ProductListView(b string) string {
	return fmt.Sprint(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Окно результатов</title>
	<style type="text/css">
		html {
			display: flex; 
			justify-content: center;
		}
		caption{
		font-weight: bold;
		margin: 20px 0px 30px 0px;
		}
		form{
			display: flex; 
			flex-direction: column;
			justify-content: center;
		}
		table {
			width: 100%;
			border-collapse: collapse;
			margin: auto;
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
		input{
		border: 0;
		width: 100%;
		color: white;
		background:deeppink;
		}
  	 </style>
    </head>
	<body>
		<button id='hello'>Hello</button>
		<table>
		<caption>Список товаров</caption>	
				<tr>
					<td>Id</td>
					<td>Наименование</td>
					<td>Колличество</td>
					<td>Цена</td>
					<td>Редактировать</td>
				</tr>  
				`, b, `
		</table>
		<br>
	<form action="http://localhost:8080/list" method="POST">
			<input  type="submit" value="Добавить">
	</form>
	<div id='trash'"></div>
</body>
<script>document.getElementById('hello').addEventListener('click', event => fetch("http://localhost:8080/list").then(data => data.text()).then(html => document.getElementById('trash').innerHTML = html))</script>
</html>
`)
}
