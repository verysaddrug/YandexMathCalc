package parse

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

// Операторные данные, включающие приоритет, ассоциативность и функцию для вычисления
var oprData = map[string]struct {
	prec  int                        // Приоритет оператора
	rAsoc bool                       // Ассоциативность: true = справа, false = слева
	fx    func(x, y float64) float64 // Функция для выполнения операции
}{
	"^": {4, true, func(x, y float64) float64 { return math.Pow(x, y) }},
	"*": {3, false, func(x, y float64) float64 { return x * y }},
	"/": {3, false, func(x, y float64) float64 { return x / y }},
	"+": {2, false, func(x, y float64) float64 { return x + y }},
	"-": {2, false, func(x, y float64) float64 { return x - y }},
}

// Функции, поддерживаемые парсером
var funcs = map[string]func(x float64) float64{
	"LN":    math.Log,
	"ABS":   math.Abs,
	"COS":   math.Cos,
	"SIN":   math.Sin,
	"TAN":   math.Tan,
	"ACOS":  math.Acos,
	"ASIN":  math.Asin,
	"ATAN":  math.Atan,
	"SQRT":  math.Sqrt,
	"CBRT":  math.Cbrt,
	"CEIL":  math.Ceil,
	"FLOOR": math.Floor,
}

// Константы, поддерживаемые парсером
var consts = map[string]float64{
	"E":       math.E,
	"PI":      math.Pi,
	"PHI":     math.Phi,
	"SQRT2":   math.Sqrt2,
	"SQRTE":   math.SqrtE,
	"SQRTPI":  math.SqrtPi,
	"SQRTPHI": math.SqrtPhi,
}

// SolvePostfix вычисляет и возвращает результат выражения, преобразованного в постфиксную форму
func SolvePostfix(tokens Stack) float64 {
	stack := Stack{}
	for _, v := range tokens.Values {
		switch v.Type {
		case NUMBER:
			stack.Push(v)
		case FUNCTION:
			stack.Push(Token{NUMBER, SolveFunction(v.Value)})
		case CONSTANT:
			if val, ok := consts[v.Value]; ok {
				stack.Push(Token{NUMBER, strconv.FormatFloat(val, 'f', -1, 64)})
			}
		case OPERATOR:
			f := oprData[v.Value].fx
			var x, y float64
			y, _ = strconv.ParseFloat(stack.Pop().Value, 64)
			x, _ = strconv.ParseFloat(stack.Pop().Value, 64)
			result := f(x, y)
			stack.Push(Token{NUMBER, strconv.FormatFloat(result, 'f', -1, 64)})
		}
	}
	out, _ := strconv.ParseFloat(stack.Values[0].Value, 64)
	return out
}

// SolveFunction возвращает результат функции, найденной в выражении
func SolveFunction(s string) string {
	var fArg float64
	fType := s[:strings.Index(s, "(")]
	args := s[strings.Index(s, "(")+1 : strings.LastIndex(s, ")")]
	if !strings.ContainsAny(args, "+ & * & - & / & ^") && !ContainsLetter(args) {
		fArg, _ = strconv.ParseFloat(args, 64)
	} else {
		stack, _ := NewParser(strings.NewReader(args)).Parse()
		stack = ShuntingYard(stack)
		fArg = SolvePostfix(stack)
	}
	return strconv.FormatFloat(funcs[fType](fArg), 'f', -1, 64)
}

// ContainsLetter проверяет, содержит ли строка буквы
func ContainsLetter(s string) bool {
	for _, v := range s {
		if unicode.IsLetter(v) {
			return true
		}
	}
	return false
}

// Solve решает выражение, представленное в виде строки
func Solve(s string) float64 {
	p := NewParser(strings.NewReader(s))
	stack, _ := p.Parse()
	stack = ShuntingYard(stack)
	answer := SolvePostfix(stack)
	return answer
}
