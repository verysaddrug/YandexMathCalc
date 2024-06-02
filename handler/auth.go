package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

// AuthMiddleware проверяет наличие JWT-токена в куки и добавляет UID в контекст запроса
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем JWT из куки
		jwtc, err := r.Cookie("jwt")
		if err != nil {
			// Перенаправляем на страницу логина, если куки не найдены
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		fmt.Println(jwtc)

		// Пример значения UID, обычно это должно быть извлечено из токена или базы данных
		v := 123

		// Копируем запрос с добавлением UID в контекст
		rcopy := r.WithContext(context.WithValue(r.Context(), "uid", v))

		// Передаем обработку следующему обработчику в цепочке
		next.ServeHTTP(w, rcopy)
	})
}

// CheckJwtMiddleware проверяет валидность JWT-токена и добавляет логин пользователя в контекст запроса
func CheckJwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем JWT из куки
		jwtc, err := r.Cookie("jwt")
		if err != nil || jwtc == nil {
			// Перенаправляем на страницу логина, если куки не найдены или пусты
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		fmt.Println(jwtc)

		const hmacSampleSecret = "super_secret_signature"

		// Парсим JWT-токен
		tokenFromString, err := jwt.Parse(jwtc.Value, func(token *jwt.Token) (interface{}, error) {
			fmt.Println("jwt.Parse")
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				// Проверяем метод подписи токена
				fmt.Println("not ok")
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Возвращаем секретный ключ для проверки подписи
			return []byte(hmacSampleSecret), nil
		})

		if err != nil {
			// Перенаправляем на страницу логина в случае ошибки парсинга токена
			fmt.Println("err not nil")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Извлекаем клеймы (claims) из токена
		claims, ok := tokenFromString.Claims.(jwt.MapClaims)
		fmt.Println("user login: ", claims["login"])

		if !ok {
			// Перенаправляем на страницу логина в случае ошибки извлечения клеймов
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Копируем запрос с добавлением логина в контекст
		rcopy := r.WithContext(context.WithValue(r.Context(), "login", claims["login"]))

		// Передаем обработку следующему обработчику в цепочке
		next.ServeHTTP(w, rcopy)
	})
}
