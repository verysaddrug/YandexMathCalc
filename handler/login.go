package handler

import (
	auth "YandexMathCalc/auth"
	db "YandexMathCalc/db"
	"fmt"
	"net/http"
	"text/template"
)

// LoginHandler обрабатывает запросы к странице логина и возвращает шаблон login.html
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим HTML-шаблон login.html
	tmpl := template.Must(template.ParseFiles("templates/login.html"))

	// Выполняем шаблон без данных (nil)
	tmpl.Execute(w, nil)
}

// LoginFormHandler обрабатывает данные формы логина и проверяет учетные данные пользователя
func LoginFormHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем логин и пароль из формы
	login := r.PostFormValue("login")
	password := r.PostFormValue("password")

	fmt.Println(login, password)

	// Проверяем пароль пользователя
	err := db.CheckPassword(login, password)
	if err != nil {
		// Если проверка не прошла, перенаправляем на страницу логина
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Создаем JWT для пользователя
	jwt := auth.CreateJwt(login)

	// Создаем куки с JWT
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: jwt,
		Path:  "/",
	}
	// Устанавливаем куки в ответе
	http.SetCookie(w, cookie)

	// Перенаправляем пользователя на главную страницу
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
