package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	calc "YandexMathCalc/calculator"
	calcp "YandexMathCalc/proto"
)

// Server представляет сервер gRPC, который реализует интерфейс CalculateServiceServer
type Server struct {
	calcp.CalculateServiceServer
}

// NewServer создает новый экземпляр сервера
func NewServer() *Server {
	return &Server{}
}

// Calculate обрабатывает запрос на вычисление выражения
func (s *Server) Calculate(ctx context.Context, req *calcp.CalculateRequest) (*calcp.CalculateResponse, error) {
	// Вызываем функцию Calculate для вычисления выражения
	res, err := calc.Calculate(req.GetExpression())
	return &calcp.CalculateResponse{
		Result: res,
	}, err
}

func main() {
	fmt.Println("starting")

	host := "localhost"
	port := "8333"

	// Формируем адрес для прослушивания
	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Println("error starting tcp listener: ", err)
		os.Exit(1)
	}

	log.Println("tcp listener started at port: ", port)

	// Создаем новый gRPC сервер
	grpcServer := grpc.NewServer()

	// Создаем новый сервер для CalculateService
	calculateServiceServer := NewServer()

	// Регистрируем сервер CalculateService на gRPC сервере
	calcp.RegisterCalculateServiceServer(grpcServer, calculateServiceServer)

	// Запускаем gRPC сервер и начинаем прослушивание подключений
	if err := grpcServer.Serve(lis); err != nil {
		log.Println("error serving grpc: ", err)
		os.Exit(1)
	}
}
