package calculator

import (
	"YandexMathCalc/parse"
	"fmt"
	"strings"
	"time"
)

// Calculate выполняет вычисление выражения, переданного в виде строки
func Calculate(s string) (float64, error) {
	// Задержка в 10 секунд для эмуляции длительных вычислений
	time.Sleep(10 * time.Second)

	// Создаем новый парсер для разбора входной строки
	p := parse.NewParser(strings.NewReader(s))

	// Парсим входную строку в стек
	stack, err := p.Parse()
	if err != nil {
		// В случае ошибки парсинга выводим сообщение об ошибке и возвращаем ошибку
		fmt.Printf("Parse error: %s\n", err)
		return 0, err
	}

	// Преобразуем инфиксное выражение в постфиксное (Обратная польская запись) с помощью алгоритма шантингового двора
	stack = parse.ShuntingYard(stack)

	// Вычисляем результат постфиксного выражения
	answer := parse.SolvePostfix(stack)
	return answer, nil
}
