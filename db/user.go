package db

import (
	mdl "YandexMathCalc/model"
	"context"
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser создает нового пользователя в базе данных
func CreateUser(name string, login string, password string) error {
	ctx := context.TODO()

	// Открываем соединение с базой данных SQLite
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Проверяем соединение с базой данных
	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	// Начинаем транзакцию
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Генерируем хэш пароля
	pwd, err := generate(password)
	if err != nil {
		return err
	}

	// Создаем нового пользователя
	user := &mdl.User{
		Name:     name,
		Login:    login,
		Password: pwd,
	}

	// Вставляем пользователя в базу данных
	userID, err := InsertUser(ctx, db, user)
	if err != nil {
		log.Println("user already exists")
		tx.Rollback()
		return err
	} else {
		user.ID = userID
	}

	// Подтверждаем транзакцию
	tx.Commit()

	return nil
}

// CheckPassword проверяет пароль пользователя
func CheckPassword(login string, password string) error {
	ctx := context.TODO()

	// Открываем соединение с базой данных SQLite
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Проверяем соединение с базой данных
	err = db.PingContext(ctx)
	if err != nil {
		return err
	}

	// Выбираем пользователя по логину
	userFromDB, err := SelectUser(ctx, db, login)
	if err != nil {
		return err
	}

	// Сравниваем хэш пароля с введенным паролем
	err = Compare(userFromDB.Password, password)
	if err != nil {
		log.Println("auth fail")
		return err
	}

	log.Println("auth success")

	return nil
}

// GetUserByLogin возвращает пользователя по его логину
func GetUserByLogin(login string) (mdl.User, error) {
	ctx := context.TODO()

	// Открываем соединение с базой данных SQLite
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return mdl.User{}, err
	}
	defer db.Close()

	// Проверяем соединение с базой данных
	err = db.PingContext(ctx)
	if err != nil {
		return mdl.User{}, err
	}

	// Выбираем пользователя по логину
	user, err := SelectUser(ctx, db, login)
	if err != nil {
		return mdl.User{}, err
	}

	return user, nil
}

// CreateUserTable создает таблицу users, если она не существует
func CreateUserTable(ctx context.Context, db *sql.DB) error {
	const usersTable = `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		login TEXT UNIQUE,
		password TEXT
	);`

	// Выполняем запрос на создание таблицы
	if _, err := db.ExecContext(ctx, usersTable); err != nil {
		return err
	}

	return nil
}

// InsertUser вставляет нового пользователя в таблицу users
func InsertUser(ctx context.Context, db *sql.DB, user *mdl.User) (int64, error) {
	var q = `
	INSERT INTO users (name, login, password) values ($1, $2, $3)
	`
	// Выполняем запрос на вставку данных
	result, err := db.ExecContext(ctx, q, user.Name, user.Login, user.Password)
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

// SelectUser выбирает пользователя по логину
func SelectUser(ctx context.Context, db *sql.DB, login string) (mdl.User, error) {
	fmt.Println("SelectUser, user login: " + login)
	var (
		user mdl.User
		err  error
	)

	var q = "SELECT id, name, login, password FROM users WHERE login=$1"
	// Выполняем запрос на выборку данных и сканируем результат в структуру пользователя
	err = db.QueryRowContext(ctx, q, login).Scan(&user.ID, &user.Name, &user.Login, &user.Password)
	return user, err
}

// Generate генерирует хэш пароля
func Generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

// Compare сравнивает хэш пароля с введенным паролем
func Compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}

// generate генерирует хэш пароля (дублирующая функция Generate)
func generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}
