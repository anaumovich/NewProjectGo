package View

import (
	"Anton/CatalogModel"
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
	button{
		border: 0;
		width: 100%;
		color: white;
		background:deeppink;
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
		arr[9] = ``
		arr[10] = `</td><td><a href="http://localhost:8080/edit?product_id=` + strconv.Itoa(i) + `"><button>Изменить</button></a></td>`
		arr[11] = `<td><a href="http://localhost:8080/delete?product_id=` + strconv.Itoa(i) + `"><button>Удалить</button></a></td>`
		arr[12] = `</tr>`

		b += strings.Join(arr, "")
	}

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
		
		<table>
		<caption>Список товаров</caption>	
				<tr>
					<td>Id</td>
					<td>Наименование</td>
					<td>Колличество</td>
					<td>Цена</td>
					<td> Процент скидки</td>
					<td>Редактировать</td>
					<td>Удалить</td>
				</tr>  
				`, b, `
		</table>
		<br>
	<form action="http://localhost:8080/list" method="POST">
			<input  type="submit" value="Добавить">
	</form>
</body>
</html>
`)
}

// todo pass by ref
func AddPageView(form CreateProductForm, headerName, buttonName string, localWay string) string {
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
    <input type="text" Name="First" placeholder="Наименование" value="`, form.Name, `">`, renderError(form.NameError), `
	<br>
    <input type="text" Name="Second" placeholder="Колличество" value="`, form.Count, `">`, renderError(form.PriceError), `
    <br>
	<input type="text" Name="Third" placeholder="Колличество" value="`, form.Price, `">`, renderError(form.PriceError), `
	<br>
	<select Name ="Выбор типа">
  		<option value = "category"> Фрукты </option>
  		<option value = "category"> Овощи </option>
		<option value = "category"> Мясо </option>
	</select> 
	<br>
	<a href="http://localhost:8080/list"><button Name = "product_id" value ="noValue">`, buttonName, `</button></a>

</form>
</body>
</html>
`)
}

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
    <input type="text" Name="First" placeholder="Наименование" value="`, form.Name, `">`, renderError(form.NameError), `
	<br>
    <input type="text" Name="Second" placeholder="Колличество" value="`, form.Count, `">`, renderError(form.CountError), `
    <br>
	<input type="text" Name="Third" placeholder="Колличество" value="`, form.Price, `">`, renderError(form.PriceError), `
	<br>
	<button Name = "product_id" value ="`+strconv.Itoa(id)+`">`, buttonName, `</button>

</form>
</body>
</html>
`)
}
