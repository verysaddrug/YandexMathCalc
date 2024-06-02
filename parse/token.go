package parse

// TokenType представляет тип токена
type TokenType int

// Token представляет токен с типом и значением
type Token struct {
	Type  TokenType // Тип токена
	Value string    // Значение токена
}

// eof представляет символ конца файла
var eof = rune(0)

// Определение типов токенов
const (
	NUMBER     TokenType = iota // Число
	LPAREN                      // Левая скобка
	RPAREN                      // Правая скобка
	CONSTANT                    // Константа
	FUNCTION                    // Функция
	OPERATOR                    // Оператор
	WHITESPACE                  // Пробел
	ERROR                       // Ошибка
	EOF                         // Конец файла
)
