package db

import (
	mdl "YandexMathCalc/model"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

// CreateExpressionsTable создает таблицу expressions, если она не существует
func CreateExpressionsTable(ctx context.Context, db *sql.DB) error {
	const expressionsTable = `
	CREATE TABLE IF NOT EXISTS expressions(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid TEXT NOT NULL,
		value TEXT NOT NULL,
		status TEXT NOT NULL,
		result TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users (id)
	);`

	// Выполняем запрос на создание таблицы
	if _, err := db.ExecContext(ctx, expressionsTable); err != nil {
		return err
	}

	return nil
}

// InsertExpression вставляет новое выражение в таблицу expressions
func InsertExpression(ctx context.Context, tx *sql.Tx, expr *mdl.Expression) (int64, error) {
	var q = `
	INSERT INTO expressions (uuid, value, status, result, user_id) values ($1, $2, $3, $4, $5)
	`

	// Выполняем запрос на вставку данных
	result, err := tx.ExecContext(ctx, q, expr.Uuid, expr.Value, expr.Status, expr.Result, expr.UserID)
	if err != nil {
		return 0, err
	}

	// Получаем ID вставленной записи
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// SelectExpressions выбирает все выражения из таблицы expressions
func SelectExpressions(ctx context.Context, db *sql.DB) ([]mdl.Expression, error) {
	var expressions []mdl.Expression
	var q = "SELECT id, uuid, value, status, result, user_id FROM expressions"

	// Выполняем запрос на выборку данных
	rows, err := db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	id := 0

	// Обрабатываем каждую строку результата
	for rows.Next() {
		e := mdl.Expression{}
		err := rows.Scan(&id, &e.Uuid, &e.Value, &e.Status, &e.Result, &e.UserID)
		if err != nil {
			return nil, err
		}
		expressions = append(expressions, e)
	}

	return expressions, nil
}

// UpdateExpression обновляет выражение в таблице expressions
func UpdateExpression(expr mdl.Expression) {
	ctx := context.TODO()

	// Открываем соединение с базой данных
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

	// Начинаем транзакцию
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	var q = `
	UPDATE expressions SET status=$1, result=$2 WHERE uuid=$3
	`

	// Выполняем запрос на обновление данных
	_, err = tx.ExecContext(ctx, q, expr.Status, expr.Result, expr.Uuid)
	if err != nil {
		fmt.Print(err)
	}

	// Подтверждаем транзакцию
	tx.Commit()
}

// GetExpressions возвращает все выражения в виде карты
func GetExpressions() map[string][]mdl.Expression {
	ctx := context.TODO()

	// Открываем соединение с базой данных
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

	// Выбираем все выражения
	exps, err := SelectExpressions(ctx, db)
	if err != nil {
		fmt.Print(err)
	}

	// Возвращаем выражения в виде карты
	r := map[string][]mdl.Expression{
		"Expressions": exps,
	}

	return r
}

// SaveExpression сохраняет новое выражение в базу данных
func SaveExpression(value string, status string, result string, login string) mdl.Expression {
	fmt.Println("SaveExpression, user login: " + login)

	id := uuid.New().String()

	ctx := context.TODO()

	// Открываем соединение с базой данных
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

	// Выбираем пользователя по логину
	user, err := SelectUser(ctx, db, login)
	if err != nil {
		panic(err)
	}

	// Создаем новое выражение
	expression := mdl.Expression{Uuid: id, Status: status, Value: value, Result: result, UserID: user.ID}

	// Начинаем транзакцию
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	// Вставляем новое выражение в базу данных
	expressionID, err := InsertExpression(ctx, tx, &expression)
	if err != nil {
		panic(err)
	}

	fmt.Println(expressionID)

	// Подтверждаем транзакцию
	tx.Commit()

	return expression
}
