package handler

import (
	db "YandexMathCalc/db"
	"fmt"
	"net/http"
	"text/template"
)

// IndexHandler обрабатывает запросы к главной странице и возвращает шаблон index.html с данными выражений
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим HTML-шаблон index.html
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// Получаем логин пользователя из контекста запроса
	login := getParams(r.Context())
	fmt.Println(login)

	// Выполняем шаблон с данными выражений, полученными из базы данных
	tmpl.Execute(w, db.GetExpressions())
}
