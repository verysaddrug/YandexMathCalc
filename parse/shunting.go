package parse

// ShuntingYard реализует алгоритм сортировочной станции для преобразования инфиксного выражения в постфиксное
func ShuntingYard(s Stack) Stack {
	postfix := Stack{}   // Стек для хранения постфиксного выражения
	operators := Stack{} // Стек для хранения операторов

	for _, v := range s.Values {
		switch v.Type {
		case OPERATOR:
			// Обработка операторов
			for !operators.IsEmpty() {
				val := v.Value
				top := operators.Peek().Value
				// Проверяем приоритет операторов и ассоциативность
				if (oprData[val].prec <= oprData[top].prec && oprData[val].rAsoc == false) ||
					(oprData[val].prec < oprData[top].prec && oprData[val].rAsoc == true) {
					// Перемещаем оператор из стека операторов в стек постфиксного выражения
					postfix.Push(operators.Pop())
					continue
				}
				break
			}
			// Помещаем текущий оператор в стек операторов
			operators.Push(v)
		case LPAREN:
			// Обработка левой скобки
			operators.Push(v)
		case RPAREN:
			// Обработка правой скобки
			for i := operators.Length() - 1; i >= 0; i-- {
				if operators.Values[i].Type != LPAREN {
					// Перемещаем операторы из стека операторов в стек постфиксного выражения до левой скобки
					postfix.Push(operators.Pop())
					continue
				} else {
					// Удаляем левую скобку из стека операторов
					operators.Pop()
					break
				}
			}
		default:
			// Все остальные токены (числа, переменные и т.д.) добавляем в стек постфиксного выражения
			postfix.Push(v)
		}
	}
	// Перемещаем все оставшиеся операторы из стека операторов в стек постфиксного выражения
	operators.EmptyInto(&postfix)
	return postfix
}
