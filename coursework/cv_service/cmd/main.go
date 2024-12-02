package main

import (
	pb "cv_service/api/proto"
	"cv_service/internal/repository"
	"cv_service/internal/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Создание нового gRPC сервера
	grpcServer := grpc.NewServer()

	storage, err := repository.New("./storage/storage.db")
	if err != nil {
		log.Fatal("failed to init database, %w", err)
	}

	// Создание нового сервера с репозиторием
	srv := server.NewServer(storage)

	// Регистрация сервиса
	pb.RegisterRecruitmentServiceServer(grpcServer, srv)

	// Установка слушателя на порту 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Запуск сервера
	log.Println("Starting gRPC server on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
