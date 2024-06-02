package parse

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"unicode"
)

// Scanner представляет сканер для разбора выражений
type Scanner struct {
	r *bufio.Reader
}

// NewScanner создает новый сканер с заданным ридером
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// Read считывает следующий символ из входного потока
func (s *Scanner) Read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof // Возвращает специальный символ eof в случае ошибки
	}
	return ch
}

// Unread возвращает последний считанный символ обратно в поток
func (s *Scanner) Unread() {
	_ = s.r.UnreadRune()
}

// Scan выполняет разбор входного потока и возвращает следующий токен
func (s *Scanner) Scan() Token {
	ch := s.Read()

	if unicode.IsDigit(ch) {
		s.Unread()
		return s.ScanNumber()
	} else if unicode.IsLetter(ch) {
		s.Unread()
		return s.ScanWord()
	} else if IsOperator(ch) {
		return Token{OPERATOR, string(ch)}
	} else if IsWhitespace(ch) {
		s.Unread()
		return s.ScanWhitespace()
	}

	switch ch {
	case eof:
		return Token{EOF, ""}
	case '(':
		return Token{LPAREN, "("}
	case ')':
		return Token{RPAREN, ")"}
	}

	return Token{ERROR, string(ch)}
}

// ScanWord считывает слово из входного потока
func (s *Scanner) ScanWord() Token {
	var buf bytes.Buffer
	buf.WriteRune(s.Read())

	for {
		if ch := s.Read(); ch == eof {
			break
		} else if ch == '(' {
			_, _ = buf.WriteRune(ch)
			parencount := 1
			for parencount > 0 {
				fch := s.Read()
				if fch == '(' {
					parencount += 1
					_, _ = buf.WriteRune(fch)
				} else if fch == ')' {
					parencount -= 1
					_, _ = buf.WriteRune(fch)
				} else {
					_, _ = buf.WriteRune(fch)
				}
			}
		} else if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			s.Unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	value := strings.ToUpper(buf.String())
	if strings.ContainsAny(value, "()") {
		return Token{FUNCTION, value}
	} else {
		return Token{CONSTANT, value}
	}
}

// ScanNumber считывает число из входного потока
func (s *Scanner) ScanNumber() Token {
	var buf bytes.Buffer
	buf.WriteRune(s.Read())

	for {
		if ch := s.Read(); ch == eof {
			break
		} else if !unicode.IsDigit(ch) && ch != '.' {
			s.Unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return Token{NUMBER, buf.String()}
}

// ScanWhitespace считывает пробелы из входного потока
func (s *Scanner) ScanWhitespace() Token {
	var buf bytes.Buffer
	buf.WriteRune(s.Read())

	for {
		if ch := s.Read(); ch == eof {
			break
		} else if !IsWhitespace(ch) {
			s.Unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return Token{WHITESPACE, buf.String()}
}

// IsOperator проверяет, является ли символ оператором
func IsOperator(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/' || r == '^'
}

// IsWhitespace проверяет, является ли символ пробелом
func IsWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}
