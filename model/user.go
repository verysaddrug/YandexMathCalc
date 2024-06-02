package model

// User представляет пользователя приложения
type User struct {
	ID              int64  // Уникальный идентификатор пользователя
	Name            string // Имя пользователя
	Login           string // Логин пользователя
	Password        string // Хэш пароля пользователя
	CurrentPassword string // Текущий пароль пользователя (может использоваться для смены пароля)
}
