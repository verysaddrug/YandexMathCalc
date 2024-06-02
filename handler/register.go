package handler

import (
	auth "YandexMathCalc/auth"
	db "YandexMathCalc/db"
	"fmt"
	"net/http"
	"text/template"
)

// RegisterHandler обрабатывает запросы к странице регистрации и возвращает шаблон register.html
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим HTML-шаблон register.html
	tmpl := template.Must(template.ParseFiles("templates/register.html"))

	// Выполняем шаблон без данных (nil)
	tmpl.Execute(w, nil)
}

// RegisterFormHandler обрабатывает данные формы регистрации и создает нового пользователя
func RegisterFormHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем значения из формы регистрации
	rn := r.PostFormValue("register-name")
	rl := r.PostFormValue("register-login")
	rp := r.PostFormValue("register-password")
	rrp := r.PostFormValue("register-repeat-password")

	fmt.Println(rn, rl, rp, rrp)

	// Проверяем, совпадают ли пароли
	if rp != rrp {
		fmt.Println("error passwords not equal")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	// Создаем нового пользователя
	err := db.CreateUser(rn, rl, rp)
	if err != nil {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	// Создаем JWT для пользователя
	jwt := auth.CreateJwt(rl)

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
