package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// CreateJwt создает JWT для указанного логина
func CreateJwt(login string) string {
	// Секретный ключ для подписи JWT
	const hmacSampleSecret = "super_secret_signature"

	// Текущее время
	now := time.Now()

	// Создание нового токена с использованием метода подписи HS256 и набора утверждений (claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,                             // Логин пользователя
		"nbf":   now.Unix(),                        // Время, с которого токен станет валидным (not before)
		"exp":   now.Add(100 * time.Minute).Unix(), // Время, после которого токен станет недействительным (expires)
		"iat":   now.Unix(),                        // Время создания токена (issued at)
	})

	// Подпись токена с использованием секретного ключа
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		panic(err) // Обработка ошибки в случае неудачи подписи
	}

	// Печать строки токена (для отладки)
	fmt.Println("token string:", tokenString)

	// Возврат строки токена
	return tokenString
}
