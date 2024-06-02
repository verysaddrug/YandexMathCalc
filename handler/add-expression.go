package handler

import (
	orch "YandexMathCalc/orchestrator"
	"fmt"
	"net/http"
	"text/template"
	"time"
)

// AddExpressionHandler обрабатывает запрос на добавление выражения
// Возвращает HTML-блок с добавленным выражением в ответ на HTMX запрос
func AddExpressionHandler(w http.ResponseWriter, r *http.Request) {
	// Эмулируем задержку в 1 секунду
	time.Sleep(1 * time.Second)

	// Получаем значение выражения из POST-запроса
	expr := r.PostFormValue("expr-val")

	// Получаем логин пользователя из контекста
	login := getParams(r.Context())

	fmt.Println("AddExpressionHandler, user login: " + login)

	// Добавляем выражение через оркестратор
	e := orch.AddExpression(expr, login)

	// Парсим HTML-шаблон и выполняем его с контекстом нового выражения
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.ExecuteTemplate(w, "expression-list-element", e)
}
