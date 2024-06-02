package db

import (
	mdl "YandexMathCalc/model"
	"fmt"
)

// InitSettings инициализирует настройки приложения
func InitSettings() {
	// Создаем пустой срез настроек
	ss := []mdl.Setting{}

	// Выводим срез в консоль
	fmt.Println(ss)
}
