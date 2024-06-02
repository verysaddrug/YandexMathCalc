package orchestrator

import (
	db "YandexMathCalc/db"
	mdl "YandexMathCalc/model"
	calcp "YandexMathCalc/proto"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Инициализация ресурсов
var resources = map[string][]mdl.ComputingResource{
	"Resources": {
		{Status: "Work", Name: "computing server", LastPing: time.Now()},
		{Status: "Reconnect", Name: "computing server", LastPing: time.Now().Add(time.Minute * (-5))},
		{Status: "Lost", Name: "computing server", LastPing: time.UnixMicro(0)},
	},
}

// GetResources возвращает текущие ресурсы
func GetResources() map[string][]mdl.ComputingResource {
	return resources
}

// AddExpression добавляет новое выражение и инициирует его вычисление
func AddExpression(expr string, login string) mdl.Expression {
	fmt.Println("AddExpression, user login: " + login)

	// Сохраняем выражение в базу данных со статусом "Calculating"
	expression := db.SaveExpression(expr, "Calculating", "?", login)

	// Запускаем вычисление выражения
	calculate(expression)

	return expression
}

// Init инициализирует оркестратор
func Init() {
	go func() {
		fmt.Println("Initialize orchestrator")
	}()
}

// calculate запускает процесс вычисления выражения в отдельной горутине
func calculate(expression mdl.Expression) {
	go func() {
		fmt.Println("calculate")

		host := "localhost"
		port := "8333"

		addr := fmt.Sprintf("%s:%s", host, port)

		// Устанавливаем соединение с gRPC сервером
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			log.Println("could not connect to grpc server: ", err)
			return
		}
		defer conn.Close()

		grpcClient := calcp.NewCalculateServiceClient(conn)

		// Отправляем запрос на вычисление выражения
		res, err := grpcClient.Calculate(context.TODO(), &calcp.CalculateRequest{Expression: expression.Value})

		if err != nil {
			expression.Status = "Error"
			expression.Result = "Error"
		} else {
			expression.Status = "Ready"
			expression.Result = fmt.Sprint(res.Result)
		}

		// Обновляем статус и результат выражения в базе данных
		db.UpdateExpression(expression)
	}()
}
