package model

import "time"

// ComputingResource представляет вычислительный ресурс
type ComputingResource struct {
	Uuid        string    // Уникальный идентификатор ресурса
	Status      string    // Статус ресурса (например, активен, неактивен)
	Name        string    // Имя ресурса
	LastPing    time.Time // Время последнего пинга ресурса
	Description string    // Описание ресурса
}
