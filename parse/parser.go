package parse

import (
	"fmt"
	"io"
)

// Parser представляет парсер для разбора выражений
type Parser struct {
	s   *Scanner // сканер для разбора токенов
	buf struct {
		tok Token // буфер для хранения токена
		n   int   // количество токенов в буфере
	}
}

// NewParser создает новый парсер с заданным ридером
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// Scan возвращает следующий токен из сканера
func (p *Parser) Scan() (tok Token) {
	// Если в буфере есть токен, возвращаем его
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok
	}

	// Иначе сканируем следующий токен
	tok = p.s.Scan()

	// Сохраняем токен в буфер
	p.buf.tok = tok

	return
}

// ScanIgnoreWhitespace возвращает следующий токен, игнорируя пробелы
func (p *Parser) ScanIgnoreWhitespace() (tok Token) {
	tok = p.Scan()
	if tok.Type == WHITESPACE {
		tok = p.Scan()
	}
	return
}

// Unscan помещает последний токен обратно в буфер
func (p *Parser) Unscan() {
	p.buf.n = 1
}

// Parse разбирает входное выражение и возвращает стек токенов
func (p *Parser) Parse() (Stack, error) {
	stack := Stack{}
	for {
		// Сканируем следующий токен, игнорируя пробелы
		tok := p.ScanIgnoreWhitespace()
		if tok.Type == ERROR {
			// Возвращаем ошибку, если токен имеет тип ERROR
			return Stack{}, fmt.Errorf("ERROR: %q", tok.Value)
		} else if tok.Type == EOF {
			// Завершаем разбор, если достигнут конец файла
			break
		} else if tok.Type == OPERATOR && tok.Value == "-" {
			// Обработка унарного минуса
			last_tok := stack.Peek()
			next_tok := p.ScanIgnoreWhitespace()
			if (last_tok.Type == OPERATOR || last_tok.Value == "" || last_tok.Type == LPAREN) && next_tok.Type == NUMBER {
				// Если предыдущий токен оператор, пустой или левая скобка, а следующий токен число
				stack.Push(Token{NUMBER, "-" + next_tok.Value})
			} else {
				// Иначе обрабатываем как бинарный минус
				stack.Push(tok)
				p.Unscan()
			}
		} else {
			// Все остальные токены добавляем в стек
			stack.Push(tok)
		}
	}
	return stack, nil
}
