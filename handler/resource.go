package handler

import (
	o "YandexMathCalc/orchestrator"
	"net/http"
	"text/template"
)

// ResourcesHandler обрабатывает запросы к странице ресурсов и возвращает шаблон resources.html с данными о ресурсах
func ResourcesHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим HTML-шаблон resources.html
	tmpl := template.Must(template.ParseFiles("templates/resources.html"))

	// Выполняем шаблон с данными о ресурсах, полученными из оркестратора
	tmpl.Execute(w, o.GetResources())
}
