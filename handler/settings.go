package handler

import (
	mdl "YandexMathCalc/model"
	"context"
	"fmt"
	"net/http"
	"text/template"
)

// SettingsHandler обрабатывает запросы к странице настроек и возвращает шаблон settings.html с данными настроек
func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим HTML-шаблон settings.html
	tmpl := template.Must(template.ParseFiles("templates/settings.html"))

	// Получаем логин пользователя из контекста запроса
	login := getParams(r.Context())
	fmt.Println(login)

	// Данные настроек, которые будут переданы в шаблон
	settings := map[string][]mdl.Setting{
		"Settings": {
			{Name: "Operation execution time +", Value: 200},
			{Name: "Operation execution time -", Value: 200},
			{Name: "Operation execution time *", Value: 200},
			{Name: "Operation execution time /", Value: 200},
			{Name: "The display time of the inactive server", Value: 200},
		},
	}

	// Выполняем шаблон с данными настроек и отправляем результат в ответ клиенту
	tmpl.Execute(w, settings)
}

// getParams извлекает значение логина пользователя из контекста
func getParams(ctx context.Context) string {
	fmt.Println("getParams")
	if ctx == nil {
		fmt.Println("getParams ctx nil")
		return ""
	}

	// Извлекаем логин пользователя из контекста
	login, ok := ctx.Value("login").(string)
	if ok {
		return login
	}

	return ""
}
