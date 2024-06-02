package db

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Init инициализирует базу данных, создавая необходимые таблицы
func Init() {
	ctx := context.TODO()

	// Открываем соединение с базой данных SQLite
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Проверяем соединение с базой данных
	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	// Создаем таблицу expressions, если она не существует
	if err = CreateExpressionsTable(ctx, db); err != nil {
		panic(err)
	}

	// Создаем таблицу users, если она не существует
	if err = CreateUserTable(ctx, db); err != nil {
		panic(err)
	}

	// Инициализируем настройки (функция не реализована)
	InitSettings()
}
