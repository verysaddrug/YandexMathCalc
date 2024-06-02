package main

import (
	db "YandexMathCalc/db"
	h "YandexMathCalc/handler"
	"fmt"
	"log"
	"net/http"
)

// go run cmd/orchestrator/main.go
func main() {
	// Выводим сообщение о запуске приложения
	fmt.Println("Distributed Calculator app...")

	// Инициализируем базу данных
	db.Init()

	// Создаем новый мультиплексор (mux) для обработки HTTP-запросов
	mux := http.NewServeMux()

	// Определяем обработчики для маршрутов, используя middleware для проверки JWT
	mux.Handle("/", h.CheckJwtMiddleware(http.HandlerFunc(h.IndexHandler)))
	mux.Handle("/add-expression", h.CheckJwtMiddleware(http.HandlerFunc(h.AddExpressionHandler)))
	mux.Handle("/settings", h.CheckJwtMiddleware(http.HandlerFunc(h.SettingsHandler)))
	mux.Handle("/resources", h.CheckJwtMiddleware(http.HandlerFunc(h.ResourcesHandler)))

	// Определяем обработчики для маршрутов аутентификации без использования middleware
	mux.HandleFunc("/login", h.LoginHandler)
	mux.HandleFunc("/register", h.RegisterHandler)

	// Определяем обработчики для форм аутентификации и регистрации
	mux.HandleFunc("/api/v1/login", h.LoginFormHandler)
	mux.HandleFunc("/api/v1/register", h.RegisterFormHandler)

	// Запускаем HTTP-сервер на порту 8000 и обрабатываем запросы с помощью mux
	if err := http.ListenAndServe(":8000", mux); err != nil {
		// В случае ошибки запуска сервера выводим сообщение об ошибке и завершаем программу
		log.Fatal(err)
	}
}
