package parse

// Stack представляет структуру данных LIFO (последним пришел - первым вышел)
type Stack struct {
	Values []Token // Срез для хранения значений токенов
}

// Pop удаляет токен с вершины стека и возвращает его значение
func (self *Stack) Pop() Token {
	if len(self.Values) == 0 {
		return Token{} // Возвращаем пустой токен, если стек пуст
	}
	token := self.Values[len(self.Values)-1]
	self.Values = self.Values[:len(self.Values)-1]
	return token
}

// Push добавляет токены на вершину стека
func (self *Stack) Push(i ...Token) {
	self.Values = append(self.Values, i...)
}

// Peek возвращает токен на вершине стека, не удаляя его
func (self *Stack) Peek() Token {
	if len(self.Values) == 0 {
		return Token{} // Возвращаем пустой токен, если стек пуст
	}
	return self.Values[len(self.Values)-1]
}

// EmptyInto перемещает все токены из одного стека в другой
func (self *Stack) EmptyInto(s *Stack) {
	if !self.IsEmpty() {
		for i := self.Length() - 1; i >= 0; i-- {
			s.Push(self.Pop())
		}
	}
}

// IsEmpty проверяет, есть ли токены в стеке
func (self *Stack) IsEmpty() bool {
	return len(self.Values) == 0
}

// Length возвращает количество токенов в стеке
func (self *Stack) Length() int {
	return len(self.Values)
}
