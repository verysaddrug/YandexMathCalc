package model

// Expression представляет математическое выражение
type Expression struct {
	Uuid   string // Уникальный идентификатор выражения
	Status string // Статус выражения (например, вычисляется, завершено)
	Value  string // Значение выражения в строковом виде
	Result string // Результат вычисления выражения
	UserID int64  // Идентификатор пользователя, которому принадлежит выражение
}