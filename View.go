package main

import (
	"fmt"
	"strconv"
)

func EditProduct(id int) string {
	return fmt.Sprint(`
<!DOCTYPE html>
<html style="display: flex; justify-content: center;">
<head>
    <meta charset="utf-8">
    <title>Кнопка отправки формы</title>
    </head>
<body>
<h3>Измените продукт </h3>

<form action="http://localhost:8080/addProduct?product_id=?" method="GET" style="display: flex; flex-direction: column;">
	<br>
    <input type="text" name="One" placeholder="Наименование" >
	<br>
    <input type="text" name="Two" placeholder="Колличество">
    <br>
	<input type="text" name="Three" placeholder="Цена">
	<br>
    <button  type="submit"  name = "product_id" value ="` + strconv.Itoa(id) + `"> Изменить</button>
</form>
</body>
</html>
`)
}

func StartPage() string {
	return fmt.Sprint(`
<!DOCTYPE html>
<html style="display: flex; justify-content: center;">
<head>
    <meta charset="utf-8">
    <title>Кнопка отправки формы</title>
    </head>
<body>
<h3>Введите новый продукт</h3>
<form action="http://localhost:8080/list" method="POST" style="display: flex; flex-direction: column;">
	<br>
    <input type="text" name="First" placeholder="Наименование" >
	<br>
    <input type="text" name="Second" placeholder="Колличество">
    <br>
	<input type="text" name="Third" placeholder="Цена">
	<br>
    <input name = "Send"  type="submit" value="Отправить">
</form>
</body>
</html>
`)
}

func NewPage(b string) string {
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
  	 </style>
    </head>
	<body>
		<table>
		<caption>Вы ввели следующие продукты</caption>	
				<tr>
					<td>Наименование</td>
					<td>Колличество</td>
					<td>Цена</td>
					<td>Редактировать</td>
				</tr>  
				`, b, `
		</table>
		<br>
	<form action="http://localhost:8080/redirect" method="POST">
			<input  type="submit" value="Добавить еще!">
	</form>
</body>
</html>
`)
}
